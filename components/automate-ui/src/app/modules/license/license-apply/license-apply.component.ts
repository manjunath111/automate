import { AfterViewInit, Component, EventEmitter, Output } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Subscription } from 'rxjs';
import { filter } from 'rxjs/operators';
import moment from 'moment/moment';

import { LicenseFacadeService, LicenseApplyReason } from '../../../entities/license/license.facade';
import { HttpStatus } from '../../../types/types';
import { ChefSessionService } from '../../../services/chef-session/chef-session.service';
import { EntityStatus, pendingState } from '../../../entities/entities';
import { ApplyStatus, FetchStatus, LicenseStatus, parsedExpirationDate } from '../../../entities/license/license.model';

@Component({
  selector: 'app-license-apply',
  templateUrl: './license-apply.component.html',
  styleUrls: ['./license-apply.component.scss']
})

export class LicenseApplyComponent implements AfterViewInit {
  @Output() openLicenseLockout = new EventEmitter<null>();

  private subscriptions: Subscription[] = [];
  public licenseApplyReason: LicenseApplyReason;
  public applyForm: FormGroup;

  public applyingLicense = false;

  // error flags
  public badRequestReason = '';
  public applyLicenseInternalError = false;
  public permissionDenied = false;
  public licenseExpired = false;

  // for rendering the success part of the modal
  public licenseApplied = false;

  // for displaying on confirmation
  public expirationDate: string;

  public modalVisible = false;
  public modalLocked = true;
  public licenseExpiredDate = null;
  public licenseType = null;

  public mlsaAgree = false;

  ngAfterViewInit(): void {
    // Set locale-aware date/time formats
    // Per https://stackoverflow.com/a/50460605, navigator.language
    // seems to be the best source of truth for browser locale.
    moment.locale(window.navigator.language);
  }

  constructor(
    public licenseFacade: LicenseFacadeService,
    private chefSessionService: ChefSessionService,
    fb: FormBuilder) {
    // keep these subscriptions always present

      this.licenseFacade.licenseApplyReason$.subscribe((reason) => {
        this.licenseApplyReason = reason;
        if (this.licenseApplyReason === LicenseApplyReason.LICENSE_ABOUT_TO_EXPIRE) {
          this.modalLocked = false;
        }
      });
      this.licenseFacade.fetchLicense$.pipe(
        filter(state => {
          let gracePeriod = false
          if (state?.license?.licensed_period) {
            this.licenseExpiredDate = moment(state?.license?.licensed_period?.end).format('ddd, DD MMM YYYY');
            this.licenseType = state?.license?.license_type;
            gracePeriod = state?.license?.grace_period;
          }
          return this.modalVisible
          && state.status === EntityStatus.loadingSuccess
          && (gracePeriod || !moment().isAfter(state.license.licensed_period.end));
        }))
        .subscribe(() => this.closeModal());

    this.applyForm = fb.group({
      licenseKey: ['', [Validators.required]]
    });
  }

  private startListening() {
    this.subscriptions.push(
      this.licenseFacade.applyLicense$.subscribe(this.handleLicenseApply.bind(this))
    );
  }

  private stopListening() {
    this.subscriptions.forEach((sub) => sub.unsubscribe());
  }

  public onApplyLicenseFormSubmit(): void {
    const license: string = this.applyForm.get('licenseKey')?.value;

    this.clearErrors();
    this.applyingLicense = true; // round-trip initiating...
    this.licenseFacade.applyLicense({ license: license.trim() });
  }

  public handleLicenseApply(state: ApplyStatus): void {
    if (pendingState(state) || !this.applyingLicense) { return; }

    this.applyingLicense = false; // round-trip complete

    if (state.status === EntityStatus.loadingSuccess) {
      this.licenseApplied = true;
      this.licenseFacade.fetchLicense$.subscribe((fetchStatus: FetchStatus) => {
        this.setExpirationDate(fetchStatus.license);
        this.unlockModal();
      });
      this.licenseFacade.getLicenseStatus();
      return;
    }

    // loadingFailure:
    switch (state.errorResp?.status) {
      case HttpStatus.FORBIDDEN:
        this.permissionDenied = true;
        break;
      case HttpStatus.BAD_REQUEST:
        const reason = state.errorResp?.error.error;
        this.badRequestReason =
          reason.includes('expired') ? 'has expired' : 'is invalid';
        break;
      default:
        this.applyLicenseInternalError = true;
    }
  }

  public openModal() {
    this.startListening();
    this.modalVisible = true;
  }

  public closeModal(): void {
    this.stopListening();
    this.modalVisible = false;
  }

  public backToLicenseLockout(): void {
    this.closeModal();
    this.openLicenseLockout.emit();
  }

  private unlockModal() {
    this.modalLocked = false;
  }

  private clearErrors(): void {
    this.applyLicenseInternalError = false;
    this.permissionDenied = false;
    this.licenseExpired = false;
    this.badRequestReason = '';
  }

  private setExpirationDate(license: LicenseStatus): void {
    this.expirationDate = parsedExpirationDate(license);
  }

  // Similarly, we must manually add this to the valid logic
  // in the submit button since chef-button doesn't play nice
  // with FormBuilder.
  public updateMLSA(event): void {
    this.mlsaAgree = event;
  }

  public logout() {
    this.chefSessionService.logout();
  }
}
