// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/queries"
	"github.com/vattle/sqlboiler/queries/qm"
	"github.com/vattle/sqlboiler/strmangle"
	"gopkg.in/nullbio/null.v6"
)

// ProductVersion is an object representing the database table.
type ProductVersion struct {
	ProductVersionID int         `boil:"product_version_id" json:"product_version_id" toml:"product_version_id" yaml:"product_version_id"`
	ProductID        int         `boil:"product_id" json:"product_id" toml:"product_id" yaml:"product_id"`
	Name             string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	GTIN             null.String `boil:"GTIN" json:"GTIN,omitempty" toml:"GTIN" yaml:"GTIN,omitempty"`
	Price            string      `boil:"price" json:"price" toml:"price" yaml:"price"`
	AddedAt          time.Time   `boil:"added_at" json:"added_at" toml:"added_at" yaml:"added_at"`
	DeletedAt        null.Time   `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	IsVisible        string      `boil:"is_visible" json:"is_visible" toml:"is_visible" yaml:"is_visible"`
	Quantity         null.String `boil:"quantity" json:"quantity,omitempty" toml:"quantity" yaml:"quantity,omitempty"`
	QuantityUnit     null.String `boil:"quantity_unit" json:"quantity_unit,omitempty" toml:"quantity_unit" yaml:"quantity_unit,omitempty"`

	R *productVersionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L productVersionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ProductVersionColumns = struct {
	ProductVersionID string
	ProductID        string
	Name             string
	GTIN             string
	Price            string
	AddedAt          string
	DeletedAt        string
	IsVisible        string
	Quantity         string
	QuantityUnit     string
}{
	ProductVersionID: "product_version_id",
	ProductID:        "product_id",
	Name:             "name",
	GTIN:             "GTIN",
	Price:            "price",
	AddedAt:          "added_at",
	DeletedAt:        "deleted_at",
	IsVisible:        "is_visible",
	Quantity:         "quantity",
	QuantityUnit:     "quantity_unit",
}

// productVersionR is where relationships are stored.
type productVersionR struct {
	Product *Product
}

// productVersionL is where Load methods for each relationship are stored.
type productVersionL struct{}

var (
	productVersionColumns               = []string{"product_version_id", "product_id", "name", "GTIN", "price", "added_at", "deleted_at", "is_visible", "quantity", "quantity_unit"}
	productVersionColumnsWithoutDefault = []string{"product_id", "name", "GTIN", "price", "deleted_at", "quantity", "quantity_unit"}
	productVersionColumnsWithDefault    = []string{"product_version_id", "added_at", "is_visible"}
	productVersionPrimaryKeyColumns     = []string{"product_version_id"}
)

type (
	// ProductVersionSlice is an alias for a slice of pointers to ProductVersion.
	// This should generally be used opposed to []ProductVersion.
	ProductVersionSlice []*ProductVersion

	productVersionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	productVersionType                 = reflect.TypeOf(&ProductVersion{})
	productVersionMapping              = queries.MakeStructMapping(productVersionType)
	productVersionPrimaryKeyMapping, _ = queries.BindMapping(productVersionType, productVersionMapping, productVersionPrimaryKeyColumns)
	productVersionInsertCacheMut       sync.RWMutex
	productVersionInsertCache          = make(map[string]insertCache)
	productVersionUpdateCacheMut       sync.RWMutex
	productVersionUpdateCache          = make(map[string]updateCache)
	productVersionUpsertCacheMut       sync.RWMutex
	productVersionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single productVersion record from the query, and panics on error.
func (q productVersionQuery) OneP() *ProductVersion {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single productVersion record from the query.
func (q productVersionQuery) One() (*ProductVersion, error) {
	o := &ProductVersion{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for product_versions")
	}

	return o, nil
}

// AllP returns all ProductVersion records from the query, and panics on error.
func (q productVersionQuery) AllP() ProductVersionSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all ProductVersion records from the query.
func (q productVersionQuery) All() (ProductVersionSlice, error) {
	var o []*ProductVersion

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ProductVersion slice")
	}

	return o, nil
}

// CountP returns the count of all ProductVersion records in the query, and panics on error.
func (q productVersionQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all ProductVersion records in the query.
func (q productVersionQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count product_versions rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q productVersionQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q productVersionQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if product_versions exists")
	}

	return count > 0, nil
}

// ProductG pointed to by the foreign key.
func (o *ProductVersion) ProductG(mods ...qm.QueryMod) productQuery {
	return o.Product(boil.GetDB(), mods...)
}

// Product pointed to by the foreign key.
func (o *ProductVersion) Product(exec boil.Executor, mods ...qm.QueryMod) productQuery {
	queryMods := []qm.QueryMod{
		qm.Where("product_id=?", o.ProductID),
	}

	queryMods = append(queryMods, mods...)

	query := Products(exec, queryMods...)
	queries.SetFrom(query.Query, "`products`")

	return query
} // LoadProduct allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (productVersionL) LoadProduct(e boil.Executor, singular bool, maybeProductVersion interface{}) error {
	var slice []*ProductVersion
	var object *ProductVersion

	count := 1
	if singular {
		object = maybeProductVersion.(*ProductVersion)
	} else {
		slice = *maybeProductVersion.(*[]*ProductVersion)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &productVersionR{}
		}
		args[0] = object.ProductID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &productVersionR{}
			}
			args[i] = obj.ProductID
		}
	}

	query := fmt.Sprintf(
		"select * from `products` where `product_id` in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Product")
	}
	defer results.Close()

	var resultSlice []*Product
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Product")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		object.R.Product = resultSlice[0]
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ProductID == foreign.ProductID {
				local.R.Product = foreign
				break
			}
		}
	}

	return nil
}

// SetProductG of the product_version to the related item.
// Sets o.R.Product to related.
// Adds o to related.R.ProductVersions.
// Uses the global database handle.
func (o *ProductVersion) SetProductG(insert bool, related *Product) error {
	return o.SetProduct(boil.GetDB(), insert, related)
}

// SetProductP of the product_version to the related item.
// Sets o.R.Product to related.
// Adds o to related.R.ProductVersions.
// Panics on error.
func (o *ProductVersion) SetProductP(exec boil.Executor, insert bool, related *Product) {
	if err := o.SetProduct(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetProductGP of the product_version to the related item.
// Sets o.R.Product to related.
// Adds o to related.R.ProductVersions.
// Uses the global database handle and panics on error.
func (o *ProductVersion) SetProductGP(insert bool, related *Product) {
	if err := o.SetProduct(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetProduct of the product_version to the related item.
// Sets o.R.Product to related.
// Adds o to related.R.ProductVersions.
func (o *ProductVersion) SetProduct(exec boil.Executor, insert bool, related *Product) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `product_versions` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"product_id"}),
		strmangle.WhereClause("`", "`", 0, productVersionPrimaryKeyColumns),
	)
	values := []interface{}{related.ProductID, o.ProductVersionID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ProductID = related.ProductID

	if o.R == nil {
		o.R = &productVersionR{
			Product: related,
		}
	} else {
		o.R.Product = related
	}

	if related.R == nil {
		related.R = &productR{
			ProductVersions: ProductVersionSlice{o},
		}
	} else {
		related.R.ProductVersions = append(related.R.ProductVersions, o)
	}

	return nil
}

// ProductVersionsG retrieves all records.
func ProductVersionsG(mods ...qm.QueryMod) productVersionQuery {
	return ProductVersions(boil.GetDB(), mods...)
}

// ProductVersions retrieves all the records using an executor.
func ProductVersions(exec boil.Executor, mods ...qm.QueryMod) productVersionQuery {
	mods = append(mods, qm.From("`product_versions`"))
	return productVersionQuery{NewQuery(exec, mods...)}
}

// FindProductVersionG retrieves a single record by ID.
func FindProductVersionG(productVersionID int, selectCols ...string) (*ProductVersion, error) {
	return FindProductVersion(boil.GetDB(), productVersionID, selectCols...)
}

// FindProductVersionGP retrieves a single record by ID, and panics on error.
func FindProductVersionGP(productVersionID int, selectCols ...string) *ProductVersion {
	retobj, err := FindProductVersion(boil.GetDB(), productVersionID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindProductVersion retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindProductVersion(exec boil.Executor, productVersionID int, selectCols ...string) (*ProductVersion, error) {
	productVersionObj := &ProductVersion{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `product_versions` where `product_version_id`=?", sel,
	)

	q := queries.Raw(exec, query, productVersionID)

	err := q.Bind(productVersionObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from product_versions")
	}

	return productVersionObj, nil
}

// FindProductVersionP retrieves a single record by ID with an executor, and panics on error.
func FindProductVersionP(exec boil.Executor, productVersionID int, selectCols ...string) *ProductVersion {
	retobj, err := FindProductVersion(exec, productVersionID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *ProductVersion) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *ProductVersion) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *ProductVersion) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *ProductVersion) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no product_versions provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(productVersionColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	productVersionInsertCacheMut.RLock()
	cache, cached := productVersionInsertCache[key]
	productVersionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			productVersionColumns,
			productVersionColumnsWithDefault,
			productVersionColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(productVersionType, productVersionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(productVersionType, productVersionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `product_versions` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `product_versions` () VALUES ()"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `product_versions` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, productVersionPrimaryKeyColumns))
		}

		if len(wl) != 0 {
			cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into product_versions")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ProductVersionID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == productVersionMapping["ProductVersionID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ProductVersionID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for product_versions")
	}

CacheNoHooks:
	if !cached {
		productVersionInsertCacheMut.Lock()
		productVersionInsertCache[key] = cache
		productVersionInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single ProductVersion record. See Update for
// whitelist behavior description.
func (o *ProductVersion) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single ProductVersion record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *ProductVersion) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the ProductVersion, and panics on error.
// See Update for whitelist behavior description.
func (o *ProductVersion) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the ProductVersion.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *ProductVersion) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	productVersionUpdateCacheMut.RLock()
	cache, cached := productVersionUpdateCache[key]
	productVersionUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			productVersionColumns,
			productVersionPrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update product_versions, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `product_versions` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, productVersionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(productVersionType, productVersionMapping, append(wl, productVersionPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update product_versions row")
	}

	if !cached {
		productVersionUpdateCacheMut.Lock()
		productVersionUpdateCache[key] = cache
		productVersionUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q productVersionQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q productVersionQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for product_versions")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ProductVersionSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o ProductVersionSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o ProductVersionSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ProductVersionSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), productVersionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `product_versions` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, productVersionPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in productVersion slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *ProductVersion) UpsertG(updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *ProductVersion) UpsertGP(updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *ProductVersion) UpsertP(exec boil.Executor, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *ProductVersion) Upsert(exec boil.Executor, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no product_versions provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(productVersionColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	productVersionUpsertCacheMut.RLock()
	cache, cached := productVersionUpsertCache[key]
	productVersionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			productVersionColumns,
			productVersionColumnsWithDefault,
			productVersionColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			productVersionColumns,
			productVersionPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert product_versions, could not build update column list")
		}

		cache.query = queries.BuildUpsertQueryMySQL(dialect, "product_versions", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `product_versions` WHERE `product_version_id`=?",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
		)

		cache.valueMapping, err = queries.BindMapping(productVersionType, productVersionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(productVersionType, productVersionMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for product_versions")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ProductVersionID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == productVersionMapping["ProductVersionID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ProductVersionID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for product_versions")
	}

CacheNoHooks:
	if !cached {
		productVersionUpsertCacheMut.Lock()
		productVersionUpsertCache[key] = cache
		productVersionUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single ProductVersion record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *ProductVersion) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single ProductVersion record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *ProductVersion) DeleteG() error {
	if o == nil {
		return errors.New("models: no ProductVersion provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single ProductVersion record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *ProductVersion) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single ProductVersion record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ProductVersion) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no ProductVersion provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), productVersionPrimaryKeyMapping)
	sql := "DELETE FROM `product_versions` WHERE `product_version_id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from product_versions")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q productVersionQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q productVersionQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no productVersionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from product_versions")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o ProductVersionSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o ProductVersionSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no ProductVersion slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o ProductVersionSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ProductVersionSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no ProductVersion slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), productVersionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `product_versions` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, productVersionPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from productVersion slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *ProductVersion) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *ProductVersion) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *ProductVersion) ReloadG() error {
	if o == nil {
		return errors.New("models: no ProductVersion provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ProductVersion) Reload(exec boil.Executor) error {
	ret, err := FindProductVersion(exec, o.ProductVersionID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *ProductVersionSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *ProductVersionSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProductVersionSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty ProductVersionSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProductVersionSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	productVersions := ProductVersionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), productVersionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `product_versions`.* FROM `product_versions` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, productVersionPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&productVersions)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ProductVersionSlice")
	}

	*o = productVersions

	return nil
}

// ProductVersionExists checks if the ProductVersion row exists.
func ProductVersionExists(exec boil.Executor, productVersionID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `product_versions` where `product_version_id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, productVersionID)
	}

	row := exec.QueryRow(sql, productVersionID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if product_versions exists")
	}

	return exists, nil
}

// ProductVersionExistsG checks if the ProductVersion row exists.
func ProductVersionExistsG(productVersionID int) (bool, error) {
	return ProductVersionExists(boil.GetDB(), productVersionID)
}

// ProductVersionExistsGP checks if the ProductVersion row exists. Panics on error.
func ProductVersionExistsGP(productVersionID int) bool {
	e, err := ProductVersionExists(boil.GetDB(), productVersionID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// ProductVersionExistsP checks if the ProductVersion row exists. Panics on error.
func ProductVersionExistsP(exec boil.Executor, productVersionID int) bool {
	e, err := ProductVersionExists(exec, productVersionID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}