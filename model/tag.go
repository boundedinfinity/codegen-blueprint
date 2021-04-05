package model

type Tag struct {
	Id          *string `json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty" validate:"required"`
	Name        *string `json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty" validate:"required"`
	Description *string `json:"description,omitempty" yaml:"description,omitempty" xml:"description,omitempty"`
}

type TagGetRequest struct {
	Id   *string `json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty" validate:"required_without=Name"`
	Name *string `json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty" validate:"required_without=Id"`
}

type TagGetResponse struct {
	Tags []Tag `json:"tags,omitempty" yaml:"tags,omitempty" xml:"tags,omitempty"`
}

type TagUpdateRequest struct {
	Tag Tag `json:"tag,omitempty" yaml:"tag,omitempty" xml:"tag,omitempty"`
}

type TagUpdateResponse struct {
	Tag Tag `json:"tag,omitempty" yaml:"tag,omitempty" xml:"tag,omitempty"`
}

type TagRemoveRequest struct {
	Id *string `json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
}

type TagRemoveResponse struct {
	Tag Tag `json:"tag,omitempty" yaml:"tag,omitempty" xml:"tag,omitempty"`
}
