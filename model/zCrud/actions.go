package zCrud

const (
	ActionList    = `list`    // retrieve rows for table
	ActionForm    = `form`    // retrieve 1 row for update
	ActionUpsert  = `upsert`  // insert if id=0, update if id>0
	ActionRestore = `restore` // same as upsert but also unset deletedAt, deletedBy
	ActionDelete  = `delete`  // same as upsert but also set deletedAt, deletedBy
)
