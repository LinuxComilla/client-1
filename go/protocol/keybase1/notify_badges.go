// Auto-generated by avdl-compiler v1.3.21 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/notify_badges.avdl

package keybase1

import (
	gregor1 "github.com/keybase/client/go/protocol/gregor1"
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type ChatConversationID []byte

func (o ChatConversationID) DeepCopy() ChatConversationID {
	return (func(x []byte) []byte {
		if x == nil {
			return nil
		}
		return append([]byte{}, x...)
	})(o)
}

type TeamMemberOutReset struct {
	Teamname string        `codec:"teamname" json:"teamname"`
	Username string        `codec:"username" json:"username"`
	Id       gregor1.MsgID `codec:"id" json:"id"`
}

func (o TeamMemberOutReset) DeepCopy() TeamMemberOutReset {
	return TeamMemberOutReset{
		Teamname: o.Teamname,
		Username: o.Username,
		Id:       o.Id.DeepCopy(),
	}
}

type BadgeState struct {
	NewTlfs                   int                     `codec:"newTlfs" json:"newTlfs"`
	RekeysNeeded              int                     `codec:"rekeysNeeded" json:"rekeysNeeded"`
	NewFollowers              int                     `codec:"newFollowers" json:"newFollowers"`
	InboxVers                 int                     `codec:"inboxVers" json:"inboxVers"`
	HomeTodoItems             int                     `codec:"homeTodoItems" json:"homeTodoItems"`
	Conversations             []BadgeConversationInfo `codec:"conversations" json:"conversations"`
	NewGitRepoGlobalUniqueIDs []string                `codec:"newGitRepoGlobalUniqueIDs" json:"newGitRepoGlobalUniqueIDs"`
	NewTeamNames              []string                `codec:"newTeamNames" json:"newTeamNames"`
	NewTeamAccessRequests     []string                `codec:"newTeamAccessRequests" json:"newTeamAccessRequests"`
	TeamsWithResetUsers       []TeamMemberOutReset    `codec:"teamsWithResetUsers" json:"teamsWithResetUsers"`
}

func (o BadgeState) DeepCopy() BadgeState {
	return BadgeState{
		NewTlfs:       o.NewTlfs,
		RekeysNeeded:  o.RekeysNeeded,
		NewFollowers:  o.NewFollowers,
		InboxVers:     o.InboxVers,
		HomeTodoItems: o.HomeTodoItems,
		Conversations: (func(x []BadgeConversationInfo) []BadgeConversationInfo {
			if x == nil {
				return nil
			}
			var ret []BadgeConversationInfo
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Conversations),
		NewGitRepoGlobalUniqueIDs: (func(x []string) []string {
			if x == nil {
				return nil
			}
			var ret []string
			for _, v := range x {
				vCopy := v
				ret = append(ret, vCopy)
			}
			return ret
		})(o.NewGitRepoGlobalUniqueIDs),
		NewTeamNames: (func(x []string) []string {
			if x == nil {
				return nil
			}
			var ret []string
			for _, v := range x {
				vCopy := v
				ret = append(ret, vCopy)
			}
			return ret
		})(o.NewTeamNames),
		NewTeamAccessRequests: (func(x []string) []string {
			if x == nil {
				return nil
			}
			var ret []string
			for _, v := range x {
				vCopy := v
				ret = append(ret, vCopy)
			}
			return ret
		})(o.NewTeamAccessRequests),
		TeamsWithResetUsers: (func(x []TeamMemberOutReset) []TeamMemberOutReset {
			if x == nil {
				return nil
			}
			var ret []TeamMemberOutReset
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.TeamsWithResetUsers),
	}
}

type BadgeConversationInfo struct {
	ConvID         ChatConversationID `codec:"convID" json:"convID"`
	BadgeCounts    map[DeviceType]int `codec:"badgeCounts" json:"badgeCounts"`
	UnreadMessages int                `codec:"unreadMessages" json:"unreadMessages"`
}

func (o BadgeConversationInfo) DeepCopy() BadgeConversationInfo {
	return BadgeConversationInfo{
		ConvID: o.ConvID.DeepCopy(),
		BadgeCounts: (func(x map[DeviceType]int) map[DeviceType]int {
			if x == nil {
				return nil
			}
			ret := make(map[DeviceType]int)
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v
				ret[kCopy] = vCopy
			}
			return ret
		})(o.BadgeCounts),
		UnreadMessages: o.UnreadMessages,
	}
}

type BadgeStateArg struct {
	BadgeState BadgeState `codec:"badgeState" json:"badgeState"`
}

type NotifyBadgesInterface interface {
	BadgeState(context.Context, BadgeState) error
}

func NotifyBadgesProtocol(i NotifyBadgesInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.NotifyBadges",
		Methods: map[string]rpc.ServeHandlerDescription{
			"badgeState": {
				MakeArg: func() interface{} {
					ret := make([]BadgeStateArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]BadgeStateArg)
					if !ok {
						err = rpc.NewTypeError((*[]BadgeStateArg)(nil), args)
						return
					}
					err = i.BadgeState(ctx, (*typedArgs)[0].BadgeState)
					return
				},
				MethodType: rpc.MethodNotify,
			},
		},
	}
}

type NotifyBadgesClient struct {
	Cli rpc.GenericClient
}

func (c NotifyBadgesClient) BadgeState(ctx context.Context, badgeState BadgeState) (err error) {
	__arg := BadgeStateArg{BadgeState: badgeState}
	err = c.Cli.Notify(ctx, "keybase.1.NotifyBadges.badgeState", []interface{}{__arg})
	return
}
