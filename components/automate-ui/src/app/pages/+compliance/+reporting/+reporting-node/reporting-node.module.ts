import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { ComplianceSharedModule } from '../../shared/shared.module';
import { ChefComponentsModule } from '../../../../components/chef-components.module';
import { ReportingNodeComponent } from './reporting-node.component';
import { ReportingNodeRoutingModule } from './reporting-node.routing';
import { CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { ChefPipesModule } from '../../../../pipes/chef-pipes.module';
import { InfiniteScrollDirective } from 'ngx-infinite-scroll';

@NgModule({
  imports: [
    CommonModule,
    RouterModule,
    ReportingNodeRoutingModule,
    ComplianceSharedModule,
    ChefComponentsModule,
    ChefPipesModule,
    InfiniteScrollDirective
  ],
  declarations: [
    ReportingNodeComponent
  ],
  schemas: [ CUSTOM_ELEMENTS_SCHEMA ],
  exports: [
    ReportingNodeComponent
  ]
})
export class ReportingNodeModule {}
