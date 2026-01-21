// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package petproject30

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/miriambudayr/pet-project-30-go/internal/apijson"
	"github.com/miriambudayr/pet-project-30-go/internal/apiquery"
	"github.com/miriambudayr/pet-project-30-go/internal/param"
	"github.com/miriambudayr/pet-project-30-go/internal/requestconfig"
	"github.com/miriambudayr/pet-project-30-go/option"
)

// PetService contains methods and other services that help with interacting with
// the pet-project-30 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPetService] method instead.
type PetService struct {
	Options []option.RequestOption
}

// NewPetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPetService(opts ...option.RequestOption) (r *PetService) {
	r = &PetService{}
	r.Options = opts
	return
}

// Creates a new pet in the store. Duplicates are allowed
func (r *PetService) New(ctx context.Context, body PetNewParams, opts ...option.RequestOption) (res *Pet, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "pets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a pet based on a single ID
func (r *PetService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *Pet, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("pets/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns all pets from the system that the user has access to Nam sed condimentum
// est. Maecenas tempor sagittis sapien, nec rhoncus sem sagittis sit amet. Aenean
// at gravida augue, ac iaculis sem. Curabitur odio lorem, ornare eget elementum
// nec, cursus id lectus. Duis mi turpis, pulvinar ac eros ac, tincidunt varius
// justo. In hac habitasse platea dictumst. Integer at adipiscing ante, a sagittis
// ligula. Aenean pharetra tempor ante molestie imperdiet. Vivamus id aliquam diam.
// Cras quis velit non tortor eleifend sagittis. Praesent at enim pharetra urna
// volutpat venenatis eget eget mauris. In eleifend fermentum facilisis. Praesent
// enim enim, gravida ac sodales sed, placerat id erat. Suspendisse lacus dolor,
// consectetur non augue vel, vehicula interdum libero. Morbi euismod sagittis
// libero sed lacinia.
//
// Sed tempus felis lobortis leo pulvinar rutrum. Nam mattis velit nisl, eu
// condimentum ligula luctus nec. Phasellus semper velit eget aliquet faucibus. In
// a mattis elit. Phasellus vel urna viverra, condimentum lorem id, rhoncus nibh.
// Ut pellentesque posuere elementum. Sed a varius odio. Morbi rhoncus ligula
// libero, vel eleifend nunc tristique vitae. Fusce et sem dui. Aenean nec
// scelerisque tortor. Fusce malesuada accumsan magna vel tempus. Quisque mollis
// felis eu dolor tristique, sit amet auctor felis gravida. Sed libero lorem,
// molestie sed nisl in, accumsan tempor nisi. Fusce sollicitudin massa ut lacinia
// mattis. Sed vel eleifend lorem. Pellentesque vitae felis pretium, pulvinar elit
// eu, euismod sapien.
func (r *PetService) List(ctx context.Context, query PetListParams, opts ...option.RequestOption) (res *[]Pet, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "pets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// deletes a single pet based on the ID supplied
func (r *PetService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("pets/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type NewPet struct {
	// Name of the pet
	Name string `json:"name,required"`
	// Type of the pet
	Tag  string     `json:"tag"`
	JSON newPetJSON `json:"-"`
}

// newPetJSON contains the JSON metadata for the struct [NewPet]
type newPetJSON struct {
	Name        apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NewPet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r newPetJSON) RawJSON() string {
	return r.raw
}

type NewPetParam struct {
	// Name of the pet
	Name param.Field[string] `json:"name,required"`
	// Type of the pet
	Tag param.Field[string] `json:"tag"`
}

func (r NewPetParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Pet struct {
	// Unique id of the pet
	ID   int64   `json:"id,required"`
	JSON petJSON `json:"-"`
	NewPet
}

// petJSON contains the JSON metadata for the struct [Pet]
type petJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Pet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petJSON) RawJSON() string {
	return r.raw
}

type PetNewParams struct {
	NewPet NewPetParam `json:"new_pet,required"`
}

func (r PetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.NewPet)
}

type PetListParams struct {
	// maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
	// tags to filter by
	Tags param.Field[[]string] `query:"tags"`
}

// URLQuery serializes [PetListParams]'s query parameters as `url.Values`.
func (r PetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
