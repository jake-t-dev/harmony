package db

type CreateTableQuery struct {
	Query string
}

type CreateDB struct {
	Queries []CreateTableQuery
}

const (
	CreateUsersTable = ` 
    CREATE TABLE IF NOT EXISTS users (
        id UUID PRIMARY KEY,
        username VARCHAR(100) NOT NULL,
        email VARCHAR(100) NOT NULL UNIQUE,
        password_hash TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`

	CreateGroupsTable = `
	CREATE TABLE IF NOT EXISTS groups (
		id UUID PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	CreateGroupInvitesTable = `
	CREATE TABLE IF NOT EXISTS group_invites (
		id UUID PRIMARY KEY,
		group_id UUID NOT NULL,
		sender_id UUID NOT NULL,
		recipient_id UUID NOT NULL,
		active BOOLEAN DEFAULT TRUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE,
		FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY (recipient_id) REFERENCES users(id) ON DELETE CASCADE
	);`

	CreateGroupMembersTable = `
	CREATE TABLE IF NOT EXISTS group_members (
		group_id UUID NOT NULL,
		user_id UUID NOT NULL,
		joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (group_id, user_id),
		FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);`

	CreateChannelsTable = `
	CREATE TABLE IF NOT EXISTS channels (
		id UUID PRIMARY KEY,
		group_id UUID NOT NULL,
		name VARCHAR(100) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE
	);`

	CreateDirectMessagesTable = `
    CREATE TABLE IF NOT EXISTS direct_messages (
        id UUID PRIMARY KEY,
        sender_id UUID NOT NULL,
        recipient_id UUID NOT NULL,
        content TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE CASCADE,
        FOREIGN KEY (recipient_id) REFERENCES users(id) ON DELETE CASCADE
    );`

	CreateGroupMessagesTable = `
	CREATE TABLE IF NOT EXISTS group_messages (
		id UUID PRIMARY KEY,
		group_id UUID NOT NULL,
		channel_id UUID NOT NULL,
		sender_id UUID NOT NULL,
		content TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE,
		FOREIGN KEY (channel_id) REFERENCES channels(id) ON DELETE CASCADE,
		FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE CASCADE
	);`

	CreateFriendsTable = `
	CREATE TABLE IF NOT EXISTS friends (
		user_id UUID NOT NULL,
		friend_id UUID NOT NULL,
		status VARCHAR(100) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (user_id, friend_id),
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY (friend_id) REFERENCES users(id) ON DELETE CASCADE
	);`
)

func GetCreateDatabase() CreateDB {

	CreateDatabase := CreateDB{
		Queries: []CreateTableQuery{
			{Query: CreateUsersTable},
			{Query: CreateGroupsTable},
			{Query: CreateGroupMembersTable},
			{Query: CreateChannelsTable},
			{Query: CreateDirectMessagesTable},
			{Query: CreateGroupMessagesTable},
			{Query: CreateFriendsTable},
		},
	}
	return CreateDatabase
}
