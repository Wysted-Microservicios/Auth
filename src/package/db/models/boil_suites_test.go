// Code generated by SQLBoiler 4.18.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Accesses", testAccesses)
	t.Run("Auths", testAuths)
	t.Run("RecoveryCodes", testRecoveryCodes)
	t.Run("RecoveryTokens", testRecoveryTokens)
	t.Run("RolesUsers", testRolesUsers)
	t.Run("Sessions", testSessions)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("Accesses", testAccessesDelete)
	t.Run("Auths", testAuthsDelete)
	t.Run("RecoveryCodes", testRecoveryCodesDelete)
	t.Run("RecoveryTokens", testRecoveryTokensDelete)
	t.Run("RolesUsers", testRolesUsersDelete)
	t.Run("Sessions", testSessionsDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Accesses", testAccessesQueryDeleteAll)
	t.Run("Auths", testAuthsQueryDeleteAll)
	t.Run("RecoveryCodes", testRecoveryCodesQueryDeleteAll)
	t.Run("RecoveryTokens", testRecoveryTokensQueryDeleteAll)
	t.Run("RolesUsers", testRolesUsersQueryDeleteAll)
	t.Run("Sessions", testSessionsQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Accesses", testAccessesSliceDeleteAll)
	t.Run("Auths", testAuthsSliceDeleteAll)
	t.Run("RecoveryCodes", testRecoveryCodesSliceDeleteAll)
	t.Run("RecoveryTokens", testRecoveryTokensSliceDeleteAll)
	t.Run("RolesUsers", testRolesUsersSliceDeleteAll)
	t.Run("Sessions", testSessionsSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Accesses", testAccessesExists)
	t.Run("Auths", testAuthsExists)
	t.Run("RecoveryCodes", testRecoveryCodesExists)
	t.Run("RecoveryTokens", testRecoveryTokensExists)
	t.Run("RolesUsers", testRolesUsersExists)
	t.Run("Sessions", testSessionsExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("Accesses", testAccessesFind)
	t.Run("Auths", testAuthsFind)
	t.Run("RecoveryCodes", testRecoveryCodesFind)
	t.Run("RecoveryTokens", testRecoveryTokensFind)
	t.Run("RolesUsers", testRolesUsersFind)
	t.Run("Sessions", testSessionsFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("Accesses", testAccessesBind)
	t.Run("Auths", testAuthsBind)
	t.Run("RecoveryCodes", testRecoveryCodesBind)
	t.Run("RecoveryTokens", testRecoveryTokensBind)
	t.Run("RolesUsers", testRolesUsersBind)
	t.Run("Sessions", testSessionsBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("Accesses", testAccessesOne)
	t.Run("Auths", testAuthsOne)
	t.Run("RecoveryCodes", testRecoveryCodesOne)
	t.Run("RecoveryTokens", testRecoveryTokensOne)
	t.Run("RolesUsers", testRolesUsersOne)
	t.Run("Sessions", testSessionsOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("Accesses", testAccessesAll)
	t.Run("Auths", testAuthsAll)
	t.Run("RecoveryCodes", testRecoveryCodesAll)
	t.Run("RecoveryTokens", testRecoveryTokensAll)
	t.Run("RolesUsers", testRolesUsersAll)
	t.Run("Sessions", testSessionsAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("Accesses", testAccessesCount)
	t.Run("Auths", testAuthsCount)
	t.Run("RecoveryCodes", testRecoveryCodesCount)
	t.Run("RecoveryTokens", testRecoveryTokensCount)
	t.Run("RolesUsers", testRolesUsersCount)
	t.Run("Sessions", testSessionsCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("Accesses", testAccessesHooks)
	t.Run("Auths", testAuthsHooks)
	t.Run("RecoveryCodes", testRecoveryCodesHooks)
	t.Run("RecoveryTokens", testRecoveryTokensHooks)
	t.Run("RolesUsers", testRolesUsersHooks)
	t.Run("Sessions", testSessionsHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Accesses", testAccessesInsert)
	t.Run("Accesses", testAccessesInsertWhitelist)
	t.Run("Auths", testAuthsInsert)
	t.Run("Auths", testAuthsInsertWhitelist)
	t.Run("RecoveryCodes", testRecoveryCodesInsert)
	t.Run("RecoveryCodes", testRecoveryCodesInsertWhitelist)
	t.Run("RecoveryTokens", testRecoveryTokensInsert)
	t.Run("RecoveryTokens", testRecoveryTokensInsertWhitelist)
	t.Run("RolesUsers", testRolesUsersInsert)
	t.Run("RolesUsers", testRolesUsersInsertWhitelist)
	t.Run("Sessions", testSessionsInsert)
	t.Run("Sessions", testSessionsInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

func TestReload(t *testing.T) {
	t.Run("Accesses", testAccessesReload)
	t.Run("Auths", testAuthsReload)
	t.Run("RecoveryCodes", testRecoveryCodesReload)
	t.Run("RecoveryTokens", testRecoveryTokensReload)
	t.Run("RolesUsers", testRolesUsersReload)
	t.Run("Sessions", testSessionsReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Accesses", testAccessesReloadAll)
	t.Run("Auths", testAuthsReloadAll)
	t.Run("RecoveryCodes", testRecoveryCodesReloadAll)
	t.Run("RecoveryTokens", testRecoveryTokensReloadAll)
	t.Run("RolesUsers", testRolesUsersReloadAll)
	t.Run("Sessions", testSessionsReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Accesses", testAccessesSelect)
	t.Run("Auths", testAuthsSelect)
	t.Run("RecoveryCodes", testRecoveryCodesSelect)
	t.Run("RecoveryTokens", testRecoveryTokensSelect)
	t.Run("RolesUsers", testRolesUsersSelect)
	t.Run("Sessions", testSessionsSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Accesses", testAccessesUpdate)
	t.Run("Auths", testAuthsUpdate)
	t.Run("RecoveryCodes", testRecoveryCodesUpdate)
	t.Run("RecoveryTokens", testRecoveryTokensUpdate)
	t.Run("RolesUsers", testRolesUsersUpdate)
	t.Run("Sessions", testSessionsUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Accesses", testAccessesSliceUpdateAll)
	t.Run("Auths", testAuthsSliceUpdateAll)
	t.Run("RecoveryCodes", testRecoveryCodesSliceUpdateAll)
	t.Run("RecoveryTokens", testRecoveryTokensSliceUpdateAll)
	t.Run("RolesUsers", testRolesUsersSliceUpdateAll)
	t.Run("Sessions", testSessionsSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
