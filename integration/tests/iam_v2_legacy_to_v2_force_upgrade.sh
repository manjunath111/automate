#!/bin/bash

# this test script:
# 1. deploys an older version of Automate and upgrades it to v2 using the beta CLI,
#    including v1 policy migration.
# 2. runs v1 inspec tests to verify the system is up and seed a v1 policy to migrate.
# 3. upgrades Automate to the latest build. This force-upgrades the system to IAM v2.
# 4. runs inspec tests to verify that the system was not disrupted by the force-upgrade
#    and legacy v1 policies continue to be enforced.

#shellcheck disable=SC2034
test_name="iam_force_upgrade_to_v2_with_legacy"
test_upgrades=true
test_upgrade_strategy="none"

# a2-deploy-integration verifies that the system is up and all APIs work correctly
# (which now includes only IAM v2 APIs)
# a2-iam-legacy-integration verifies that migrated v1 legacy policies persisted 
# and their permissions are enforced
test_upgrade_inspec_profiles=(a2-deploy-integration a2-iam-legacy-integration)

# Note: we can't run diagnostics AND inspec, so skip diagnostics
test_skip_diagnostics=true

# on this version, Automate had upgrade-to-v2 and the first three v2 data migrations
OLD_VERSION=20190501153509
OLD_MANIFEST_DIR="${A2_ROOT_DIR}/components/automate-deployment/testdata/old_manifests/"
DEEP_UPGRADE_PATH="${OLD_MANIFEST_DIR}/${OLD_VERSION}.json"

do_deploy() {
    #shellcheck disable=SC2154
    cp "$DEEP_UPGRADE_PATH" "$test_manifest_path"

    # we make sure to use the CLI for the old version of Automate we want to deploy
    local cli_bin="/bin/chef-automate-${OLD_VERSION}"

    download_cli "${OLD_VERSION}" "${cli_bin}"

    #shellcheck disable=SC2154
    "${cli_bin}" deploy "$test_config_path" \
        --hartifacts "$test_hartifacts_path" \
        --override-origin "$HAB_ORIGIN" \
        --manifest-dir "$test_manifest_path" \
        --admin-password chefautomate \
        --accept-terms-and-mlsa \
        --skip-preflight \
        --debug

    # one by-product of these tests is a custom v1 policy
    # it should be:
    # - migrated to v2 by the upgrade-to-v2 command below
    # - preserved by the force-uprade
    # they also verify that the system deployed successfully
    run_inspec_tests "${A2_ROOT_DIR}" "a2-iam-v1-integration"

    # this migrates v1 legacy policies
    "${cli_bin}" iam upgrade-to-v2 

  do_apply_license
}

do_prepare_upgrade() {
  # use latest current here
  prepare_upgrade_milestone "current" "20220329091442"
}
