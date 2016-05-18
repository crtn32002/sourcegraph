// GENERATED CODE - DO NOT EDIT!
// @generated
//
// Generated by:
//
//   go run gen_context_and_mock.go -o1 context.go -o2 mockstore/mockstores.go
//
// Called via:
//
//   go generate
//

package store

import (
	"golang.org/x/net/context"
	srcstore "sourcegraph.com/sourcegraph/srclib/store"
)

// Stores has a field for each store interface.
type Stores struct {
	Accounts           Accounts
	BuildLogs          BuildLogs
	Builds             Builds
	Channel            Channel
	Directory          Directory
	ExternalAuthTokens ExternalAuthTokens
	GlobalDefs         GlobalDefs
	GlobalDeps         GlobalDeps
	GlobalRefs         GlobalRefs
	Graph              srcstore.MultiRepoStoreImporterIndexer
	Orgs               Orgs
	Password           Password
	RepoConfigs        RepoConfigs
	RepoPerms          RepoPerms
	RepoStatuses       RepoStatuses
	RepoVCS            RepoVCS
	Repos              Repos
	Users              Users
}

type contextKey int

const (
	_AccountsKey contextKey = iota
	_BuildLogsKey
	_BuildsKey
	_ChannelKey
	_DirectoryKey
	_ExternalAuthTokensKey
	_GlobalDefsKey
	_GlobalDepsKey
	_GlobalRefsKey
	_GraphKey
	_OrgsKey
	_PasswordKey
	_RepoConfigsKey
	_RepoPermsKey
	_RepoStatusesKey
	_RepoVCSKey
	_ReposKey
	_UsersKey
)

// WithStores returns a copy of parent with the given stores. If a store's field value is nil, its previous value is inherited from parent in the new context.
func WithStores(ctx context.Context, s Stores) context.Context {
	if s.Accounts != nil {
		ctx = WithAccounts(ctx, s.Accounts)
	}
	if s.BuildLogs != nil {
		ctx = WithBuildLogs(ctx, s.BuildLogs)
	}
	if s.Builds != nil {
		ctx = WithBuilds(ctx, s.Builds)
	}
	if s.Channel != nil {
		ctx = WithChannel(ctx, s.Channel)
	}
	if s.Directory != nil {
		ctx = WithDirectory(ctx, s.Directory)
	}
	if s.ExternalAuthTokens != nil {
		ctx = WithExternalAuthTokens(ctx, s.ExternalAuthTokens)
	}
	if s.GlobalDefs != nil {
		ctx = WithGlobalDefs(ctx, s.GlobalDefs)
	}
	if s.GlobalDeps != nil {
		ctx = WithGlobalDeps(ctx, s.GlobalDeps)
	}
	if s.GlobalRefs != nil {
		ctx = WithGlobalRefs(ctx, s.GlobalRefs)
	}
	if s.Graph != nil {
		ctx = WithGraph(ctx, s.Graph)
	}
	if s.Orgs != nil {
		ctx = WithOrgs(ctx, s.Orgs)
	}
	if s.Password != nil {
		ctx = WithPassword(ctx, s.Password)
	}
	if s.RepoConfigs != nil {
		ctx = WithRepoConfigs(ctx, s.RepoConfigs)
	}
	if s.RepoPerms != nil {
		ctx = WithRepoPerms(ctx, s.RepoPerms)
	}
	if s.RepoStatuses != nil {
		ctx = WithRepoStatuses(ctx, s.RepoStatuses)
	}
	if s.RepoVCS != nil {
		ctx = WithRepoVCS(ctx, s.RepoVCS)
	}
	if s.Repos != nil {
		ctx = WithRepos(ctx, s.Repos)
	}
	if s.Users != nil {
		ctx = WithUsers(ctx, s.Users)
	}
	return ctx
}

// WithAccounts returns a copy of parent with the given Accounts store.
func WithAccounts(parent context.Context, s Accounts) context.Context {
	return context.WithValue(parent, _AccountsKey, s)
}

// AccountsFromContext gets the context's Accounts store. If the store is not present, it panics.
func AccountsFromContext(ctx context.Context) Accounts {
	s, ok := ctx.Value(_AccountsKey).(Accounts)
	if !ok || s == nil {
		panic("no Accounts set in context")
	}
	return s
}

// WithBuildLogs returns a copy of parent with the given BuildLogs store.
func WithBuildLogs(parent context.Context, s BuildLogs) context.Context {
	return context.WithValue(parent, _BuildLogsKey, s)
}

// BuildLogsFromContext gets the context's BuildLogs store. If the store is not present, it panics.
func BuildLogsFromContext(ctx context.Context) BuildLogs {
	s, ok := ctx.Value(_BuildLogsKey).(BuildLogs)
	if !ok || s == nil {
		panic("no BuildLogs set in context")
	}
	return s
}

// WithBuilds returns a copy of parent with the given Builds store.
func WithBuilds(parent context.Context, s Builds) context.Context {
	return context.WithValue(parent, _BuildsKey, s)
}

// BuildsFromContext gets the context's Builds store. If the store is not present, it panics.
func BuildsFromContext(ctx context.Context) Builds {
	s, ok := ctx.Value(_BuildsKey).(Builds)
	if !ok || s == nil {
		panic("no Builds set in context")
	}
	return s
}

// WithChannel returns a copy of parent with the given Channel store.
func WithChannel(parent context.Context, s Channel) context.Context {
	return context.WithValue(parent, _ChannelKey, s)
}

// ChannelFromContext gets the context's Channel store. If the store is not present, it panics.
func ChannelFromContext(ctx context.Context) Channel {
	s, ok := ctx.Value(_ChannelKey).(Channel)
	if !ok || s == nil {
		panic("no Channel set in context")
	}
	return s
}

// WithDirectory returns a copy of parent with the given Directory store.
func WithDirectory(parent context.Context, s Directory) context.Context {
	return context.WithValue(parent, _DirectoryKey, s)
}

// DirectoryFromContext gets the context's Directory store. If the store is not present, it panics.
func DirectoryFromContext(ctx context.Context) Directory {
	s, ok := ctx.Value(_DirectoryKey).(Directory)
	if !ok || s == nil {
		panic("no Directory set in context")
	}
	return s
}

// WithExternalAuthTokens returns a copy of parent with the given ExternalAuthTokens store.
func WithExternalAuthTokens(parent context.Context, s ExternalAuthTokens) context.Context {
	return context.WithValue(parent, _ExternalAuthTokensKey, s)
}

// ExternalAuthTokensFromContext gets the context's ExternalAuthTokens store. If the store is not present, it panics.
func ExternalAuthTokensFromContext(ctx context.Context) ExternalAuthTokens {
	s, ok := ctx.Value(_ExternalAuthTokensKey).(ExternalAuthTokens)
	if !ok || s == nil {
		panic("no ExternalAuthTokens set in context")
	}
	return s
}

// WithGlobalDefs returns a copy of parent with the given GlobalDefs store.
func WithGlobalDefs(parent context.Context, s GlobalDefs) context.Context {
	return context.WithValue(parent, _GlobalDefsKey, s)
}

// GlobalDefsFromContext gets the context's GlobalDefs store. If the store is not present, it panics.
func GlobalDefsFromContext(ctx context.Context) GlobalDefs {
	s, ok := ctx.Value(_GlobalDefsKey).(GlobalDefs)
	if !ok || s == nil {
		panic("no GlobalDefs set in context")
	}
	return s
}

// WithGlobalDeps returns a copy of parent with the given GlobalDeps store.
func WithGlobalDeps(parent context.Context, s GlobalDeps) context.Context {
	return context.WithValue(parent, _GlobalDepsKey, s)
}

// GlobalDepsFromContext gets the context's GlobalDeps store. If the store is not present, it panics.
func GlobalDepsFromContext(ctx context.Context) GlobalDeps {
	s, ok := ctx.Value(_GlobalDepsKey).(GlobalDeps)
	if !ok || s == nil {
		panic("no GlobalDeps set in context")
	}
	return s
}

// WithGlobalRefs returns a copy of parent with the given GlobalRefs store.
func WithGlobalRefs(parent context.Context, s GlobalRefs) context.Context {
	return context.WithValue(parent, _GlobalRefsKey, s)
}

// GlobalRefsFromContext gets the context's GlobalRefs store. If the store is not present, it panics.
func GlobalRefsFromContext(ctx context.Context) GlobalRefs {
	s, ok := ctx.Value(_GlobalRefsKey).(GlobalRefs)
	if !ok || s == nil {
		panic("no GlobalRefs set in context")
	}
	return s
}

// WithGraph returns a copy of parent with the given Graph store.
func WithGraph(parent context.Context, s srcstore.MultiRepoStoreImporterIndexer) context.Context {
	return context.WithValue(parent, _GraphKey, s)
}

// GraphFromContext gets the context's Graph store. If the store is not present, it panics.
func GraphFromContext(ctx context.Context) srcstore.MultiRepoStoreImporterIndexer {
	s, ok := ctx.Value(_GraphKey).(srcstore.MultiRepoStoreImporterIndexer)
	if !ok || s == nil {
		panic("no Graph set in context")
	}
	return s
}

// WithOrgs returns a copy of parent with the given Orgs store.
func WithOrgs(parent context.Context, s Orgs) context.Context {
	return context.WithValue(parent, _OrgsKey, s)
}

// OrgsFromContext gets the context's Orgs store. If the store is not present, it panics.
func OrgsFromContext(ctx context.Context) Orgs {
	s, ok := ctx.Value(_OrgsKey).(Orgs)
	if !ok || s == nil {
		panic("no Orgs set in context")
	}
	return s
}

// WithPassword returns a copy of parent with the given Password store.
func WithPassword(parent context.Context, s Password) context.Context {
	return context.WithValue(parent, _PasswordKey, s)
}

// PasswordFromContext gets the context's Password store. If the store is not present, it panics.
func PasswordFromContext(ctx context.Context) Password {
	s, ok := ctx.Value(_PasswordKey).(Password)
	if !ok || s == nil {
		panic("no Password set in context")
	}
	return s
}

// WithRepoConfigs returns a copy of parent with the given RepoConfigs store.
func WithRepoConfigs(parent context.Context, s RepoConfigs) context.Context {
	return context.WithValue(parent, _RepoConfigsKey, s)
}

// RepoConfigsFromContext gets the context's RepoConfigs store. If the store is not present, it panics.
func RepoConfigsFromContext(ctx context.Context) RepoConfigs {
	s, ok := ctx.Value(_RepoConfigsKey).(RepoConfigs)
	if !ok || s == nil {
		panic("no RepoConfigs set in context")
	}
	return s
}

// WithRepoPerms returns a copy of parent with the given RepoPerms store.
func WithRepoPerms(parent context.Context, s RepoPerms) context.Context {
	return context.WithValue(parent, _RepoPermsKey, s)
}

// RepoPermsFromContext gets the context's RepoPerms store. If the store is not present, it panics.
func RepoPermsFromContext(ctx context.Context) RepoPerms {
	s, ok := ctx.Value(_RepoPermsKey).(RepoPerms)
	if !ok || s == nil {
		panic("no RepoPerms set in context")
	}
	return s
}

// WithRepoStatuses returns a copy of parent with the given RepoStatuses store.
func WithRepoStatuses(parent context.Context, s RepoStatuses) context.Context {
	return context.WithValue(parent, _RepoStatusesKey, s)
}

// RepoStatusesFromContext gets the context's RepoStatuses store. If the store is not present, it panics.
func RepoStatusesFromContext(ctx context.Context) RepoStatuses {
	s, ok := ctx.Value(_RepoStatusesKey).(RepoStatuses)
	if !ok || s == nil {
		panic("no RepoStatuses set in context")
	}
	return s
}

// WithRepoVCS returns a copy of parent with the given RepoVCS store.
func WithRepoVCS(parent context.Context, s RepoVCS) context.Context {
	return context.WithValue(parent, _RepoVCSKey, s)
}

// RepoVCSFromContext gets the context's RepoVCS store. If the store is not present, it panics.
func RepoVCSFromContext(ctx context.Context) RepoVCS {
	s, ok := ctx.Value(_RepoVCSKey).(RepoVCS)
	if !ok || s == nil {
		panic("no RepoVCS set in context")
	}
	return s
}

// WithRepos returns a copy of parent with the given Repos store.
func WithRepos(parent context.Context, s Repos) context.Context {
	return context.WithValue(parent, _ReposKey, s)
}

// ReposFromContext gets the context's Repos store. If the store is not present, it panics.
func ReposFromContext(ctx context.Context) Repos {
	s, ok := ctx.Value(_ReposKey).(Repos)
	if !ok || s == nil {
		panic("no Repos set in context")
	}
	return s
}

// WithUsers returns a copy of parent with the given Users store.
func WithUsers(parent context.Context, s Users) context.Context {
	return context.WithValue(parent, _UsersKey, s)
}

// UsersFromContext gets the context's Users store. If the store is not present, it panics.
func UsersFromContext(ctx context.Context) Users {
	s, ok := ctx.Value(_UsersKey).(Users)
	if !ok || s == nil {
		panic("no Users set in context")
	}
	return s
}
