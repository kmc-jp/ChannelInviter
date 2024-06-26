# 邀請機器人

這個軟件是用來邀請其他人進入隱藏頻道的。通過使用這個軟件，我們可以隨意進入已登記的隱藏頻道。

## 安裝軟件

首先我們需要得到SlackAPI Token. 這個是為了做Slack App的App Manifest.

```json
{
    "display_information": {
        "name": "Channel Inviter",
        "description": "Channel Inviter invites users to pre-specified channels",
        "background_color": "#114a80"
    },
    "features": {
        "bot_user": {
            "display_name": "Channel Inviter",
            "always_online": false
        },
        "slash_commands": [
            {
                "command": "/inviterinvite",
                "description": "Invite user to channels registered to the Keyword",
                "usage_hint": "Keyword User1 User2...",
                "should_escape": false
            },
            {
                "command": "/inviterjoin",
                "description": "Join channels registered to the Keyword",
                "usage_hint": "Keyword",
                "should_escape": false
            },
            {
                "command": "/inviteraddchannels",
                "description": "Register channels to the keyword",
                "usage_hint": "Keyword #Channel1 #Channel2...",
                "should_escape": false
            },
            {
                "command": "/inviterremovechannels",
                "description": "Remove channels registered to the keyword",
                "usage_hint": "Keyword #Channel1 #Channel2...",
                "should_escape": false
            }
        ]
    },
    "oauth_config": {
        "scopes": {
            "bot": [
                "app_mentions:read",
                "channels:manage",
                "channels:write.invites",
                "groups:write",
                "groups:write.invites",
                "im:write",
                "mpim:write",
                "chat:write",
                "commands"
            ]
        }
    },
    "settings": {
        "event_subscriptions": {
            "bot_events": [
                "app_mention"
            ]
        },
        "interactivity": {
            "is_enabled": true
        },
        "org_deploy_enabled": false,
        "socket_mode_enabled": true,
        "token_rotation_enabled": false
    }
}
```

為了使用Socket mode, App Level Token也你必須得到.

然後, 請下載此儲存庫, 修改 `settings/settings.yaml`.
