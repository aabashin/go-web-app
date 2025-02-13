package data

import (
    up "github.com/upper/db/v4"
    "time"
)
// Testmodel struct
type Testmodel struct {
    ID        int       `db:"id,omitempty"`
    CreatedAt time.Time `db:"created_at"`
    UpdatedAt time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *Testmodel) Table() string {
    return "testmodels"
}

// GetAll gets all records from the database, using upper
func (t *Testmodel) GetAll(condition up.Cond) ([]*Testmodel, error) {
    collection := upper.Collection(t.Table())
    var all []*Testmodel

    res := collection.Find(condition)
    err := res.All(&all)
    if err != nil {
        return nil, err
    }

    return all, err
}

// Get gets one record from the database, by id, using upper
func (t *Testmodel) Get(id int) (*Testmodel, error) {
    var one Testmodel
    collection := upper.Collection(t.Table())

    res := collection.Find(up.Cond{"id": id})
    err := res.One(&one)
    if err != nil {
        return nil, err
    }
    return &one, nil
}

// Update updates a record in the database, using upper
func (t *Testmodel) Update(m Testmodel) error {
    m.UpdatedAt = time.Now()
    collection := upper.Collection(t.Table())
    res := collection.Find(m.ID)
    err := res.Update(&m)
    if err != nil {
        return err
    }
    return nil
}

// Delete deletes a record from the database by id, using upper
func (t *Testmodel) Delete(id int) error {
    collection := upper.Collection(t.Table())
    res := collection.Find(id)
    err := res.Delete()
    if err != nil {
        return err
    }
    return nil
}

// Insert inserts a model into the database, using upper
func (t *Testmodel) Insert(m Testmodel) (int, error) {
    m.CreatedAt = time.Now()
    m.UpdatedAt = time.Now()
    collection := upper.Collection(t.Table())
    res, err := collection.Insert(m)
    if err != nil {
        return 0, err
    }

    id := getInsertID(res.ID())

    return id, nil
}

// Builder is an example of using upper's sql builder
func (t *Testmodel) Builder(id int) ([]*Testmodel, error) {
    collection := upper.Collection(t.Table())

    var result []*Testmodel

    err := collection.Session().
        SQL().
        SelectFrom(t.Table()).
        Where("id > ?", id).
        OrderBy("id").
        All(&result)
    if err != nil {
        return nil, err
    }
    return result, nil
}

