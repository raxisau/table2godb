# table2godb

Dumps out the Structure that is used to describe a table that can be used in godb - <https://github.com/samonzeweb/godb>

Table2godb is a **prototype** for creating a Go [struct](https://golang.org/ref/spec#Struct_types) from an database table. The resulting struct works  for [godb - a Go query builder and struct mapper](https://github.com/samonzeweb/godb).

This was inspired by [Zek](https://github.com/miku/zek) to shorten the time to create Data Access Structures for databases.

```bash
table2godb "-ds=root:password@tcp(127.0.0.1:3306)/mysql" -t=user
```

```go
type User struct {
    Host                   string    `db:"Host,key"`
    User                   string    `db:"User,key"`
    SelectPriv             string    `db:"Select_priv"`
    InsertPriv             string    `db:"Insert_priv"`
    UpdatePriv             string    `db:"Update_priv"`
    DeletePriv             string    `db:"Delete_priv"`
    CreatePriv             string    `db:"Create_priv"`
    DropPriv               string    `db:"Drop_priv"`
    ReloadPriv             string    `db:"Reload_priv"`
    ShutdownPriv           string    `db:"Shutdown_priv"`
    ProcessPriv            string    `db:"Process_priv"`
    FilePriv               string    `db:"File_priv"`
    GrantPriv              string    `db:"Grant_priv"`
    ReferencesPriv         string    `db:"References_priv"`
    IndexPriv              string    `db:"Index_priv"`
    AlterPriv              string    `db:"Alter_priv"`
    ShowDbPriv             string    `db:"Show_db_priv"`
    SuperPriv              string    `db:"Super_priv"`
    CreateTmpTablePriv     string    `db:"Create_tmp_table_priv"`
    LockTablesPriv         string    `db:"Lock_tables_priv"`
    ExecutePriv            string    `db:"Execute_priv"`
    ReplSlavePriv          string    `db:"Repl_slave_priv"`
    ReplClientPriv         string    `db:"Repl_client_priv"`
    CreateViewPriv         string    `db:"Create_view_priv"`
    ShowViewPriv           string    `db:"Show_view_priv"`
    CreateRoutinePriv      string    `db:"Create_routine_priv"`
    AlterRoutinePriv       string    `db:"Alter_routine_priv"`
    CreateUserPriv         string    `db:"Create_user_priv"`
    EventPriv              string    `db:"Event_priv"`
    TriggerPriv            string    `db:"Trigger_priv"`
    CreateTablespacePriv   string    `db:"Create_tablespace_priv"`
    SslType                string    `db:"ssl_type"`
    SslCipher              string    `db:"ssl_cipher"`
    X509Issuer             string    `db:"x509_issuer"`
    X509Subject            string    `db:"x509_subject"`
    MaxQuestions           int64     `db:"max_questions"`
    MaxUpdates             int64     `db:"max_updates"`
    MaxConnections         int64     `db:"max_connections"`
    MaxUserConnections     int64     `db:"max_user_connections"`
    Plugin                 string    `db:"plugin"`
    AuthenticationString   string    `db:"authentication_string"`
    PasswordExpired        string    `db:"password_expired"`
    PasswordLastChanged    time.Time `db:"password_last_changed"`
    PasswordLifetime       int64     `db:"password_lifetime"`
    AccountLocked          string    `db:"account_locked"`
    CreateRolePriv         string    `db:"Create_role_priv"`
    DropRolePriv           string    `db:"Drop_role_priv"`
    PasswordReuseHistory   int64     `db:"Password_reuse_history"`
    PasswordReuseTime      int64     `db:"Password_reuse_time"`
    PasswordRequireCurrent string    `db:"Password_require_current"`
    UserAttributes         string    `db:"User_attributes"`
}

// TableName name of the table
func (*User) TableName() string {
    return "user"
}
```

## Install

```bash
go get github.com/raxisau/table2godb/cmd...
```

## Usage

```shell
$ table2godb -h
Usage of table2godb:
  -ds string
        Data source for the database (default "username:password@tcp(127.0.0.1:3306)/mysql")
  -t string
        Name of the table to code up (default "user")
```
