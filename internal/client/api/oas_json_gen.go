// Code generated by ogen, DO NOT EDIT.

package api

import (
	"math/bits"
	"strconv"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/validate"
)

// Encode implements json.Marshaler.
func (s *Directory) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Directory) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("root")
		e.Int(s.Root)
	}
	{
		e.FieldStart("relative_path")
		e.Str(s.RelativePath)
	}
	{
		e.FieldStart("content")
		e.ArrStart()
		for _, elem := range s.Content {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
}

var jsonFieldsNameOfDirectory = [3]string{
	0: "root",
	1: "relative_path",
	2: "content",
}

// Decode decodes Directory from json.
func (s *Directory) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Directory to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "root":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int()
				s.Root = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"root\"")
			}
		case "relative_path":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.RelativePath = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"relative_path\"")
			}
		case "content":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				s.Content = make([]File, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem File
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Content = append(s.Content, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"content\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Directory")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfDirectory) {
					name = jsonFieldsNameOfDirectory[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Directory) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Directory) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *File) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *File) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("name")
		e.Str(s.Name)
	}
	{
		e.FieldStart("directory")
		e.Bool(s.Directory)
	}
}

var jsonFieldsNameOfFile = [2]string{
	0: "name",
	1: "directory",
}

// Decode decodes File from json.
func (s *File) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode File to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "name":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Name = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		case "directory":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Bool()
				s.Directory = bool(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"directory\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode File")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfFile) {
					name = jsonFieldsNameOfFile[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *File) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *File) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes GetStatusOK as json.
func (s *GetStatusOK) Encode(e *jx.Encoder) {
	unwrapped := (*Status)(s)

	unwrapped.Encode(e)
}

// Decode decodes GetStatusOK from json.
func (s *GetStatusOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GetStatusOK to nil")
	}
	var unwrapped Status
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = GetStatusOK(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GetStatusOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GetStatusOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes GetStatusServiceUnavailable as json.
func (s *GetStatusServiceUnavailable) Encode(e *jx.Encoder) {
	unwrapped := (*Status)(s)

	unwrapped.Encode(e)
}

// Decode decodes GetStatusServiceUnavailable from json.
func (s *GetStatusServiceUnavailable) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GetStatusServiceUnavailable to nil")
	}
	var unwrapped Status
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = GetStatusServiceUnavailable(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GetStatusServiceUnavailable) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GetStatusServiceUnavailable) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes bool as json.
func (o OptBool) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Bool(bool(o.Value))
}

// Decode decodes bool from json.
func (o *OptBool) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptBool to nil")
	}
	o.Set = true
	v, err := d.Bool()
	if err != nil {
		return err
	}
	o.Value = bool(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptBool) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptBool) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes PlaybackPosition as json.
func (o OptPlaybackPosition) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes PlaybackPosition from json.
func (o *OptPlaybackPosition) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptPlaybackPosition to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptPlaybackPosition) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptPlaybackPosition) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes PutPlayerPauseReq as json.
func (o OptPutPlayerPauseReq) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes PutPlayerPauseReq from json.
func (o *OptPutPlayerPauseReq) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptPutPlayerPauseReq to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptPutPlayerPauseReq) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptPutPlayerPauseReq) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes string as json.
func (o OptString) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptString to nil")
	}
	o.Set = true
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Pause) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Pause) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("paused")
		e.Bool(s.Paused)
	}
	{
		if s.PausedMinusAt.Set {
			e.FieldStart("paused-at")
			s.PausedMinusAt.Encode(e)
		}
	}
}

var jsonFieldsNameOfPause = [2]string{
	0: "paused",
	1: "paused-at",
}

// Decode decodes Pause from json.
func (s *Pause) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Pause to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "paused":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Bool()
				s.Paused = bool(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"paused\"")
			}
		case "paused-at":
			if err := func() error {
				s.PausedMinusAt.Reset()
				if err := s.PausedMinusAt.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"paused-at\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Pause")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfPause) {
					name = jsonFieldsNameOfPause[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Pause) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Pause) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes PlaybackPosition as json.
func (s PlaybackPosition) Encode(e *jx.Encoder) {
	unwrapped := float32(s)

	e.Float32(unwrapped)
}

// Decode decodes PlaybackPosition from json.
func (s *PlaybackPosition) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode PlaybackPosition to nil")
	}
	var unwrapped float32
	if err := func() error {
		v, err := d.Float32()
		unwrapped = float32(v)
		if err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = PlaybackPosition(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s PlaybackPosition) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *PlaybackPosition) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *PutPlayerPauseReq) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *PutPlayerPauseReq) encodeFields(e *jx.Encoder) {
	{
		if s.Pause.Set {
			e.FieldStart("pause")
			s.Pause.Encode(e)
		}
	}
}

var jsonFieldsNameOfPutPlayerPauseReq = [1]string{
	0: "pause",
}

// Decode decodes PutPlayerPauseReq from json.
func (s *PutPlayerPauseReq) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode PutPlayerPauseReq to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "pause":
			if err := func() error {
				s.Pause.Reset()
				if err := s.Pause.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"pause\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode PutPlayerPauseReq")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *PutPlayerPauseReq) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *PutPlayerPauseReq) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Status) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Status) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("connection")
		s.Connection.Encode(e)
	}
	{
		e.FieldStart("player")
		e.Str(s.Player)
	}
	{
		e.FieldStart("version")
		s.Version.Encode(e)
	}
}

var jsonFieldsNameOfStatus = [3]string{
	0: "connection",
	1: "player",
	2: "version",
}

// Decode decodes Status from json.
func (s *Status) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Status to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "connection":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				if err := s.Connection.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"connection\"")
			}
		case "player":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Player = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"player\"")
			}
		case "version":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				if err := s.Version.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"version\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Status")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfStatus) {
					name = jsonFieldsNameOfStatus[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Status) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Status) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes StatusConnection as json.
func (s StatusConnection) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

// Decode decodes StatusConnection from json.
func (s *StatusConnection) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode StatusConnection to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	// Try to use constant string.
	switch StatusConnection(v) {
	case StatusConnectionOk:
		*s = StatusConnectionOk
	case StatusConnectionConnecting:
		*s = StatusConnectionConnecting
	case StatusConnectionDisconnected:
		*s = StatusConnectionDisconnected
	default:
		*s = StatusConnection(v)
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s StatusConnection) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *StatusConnection) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Version) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Version) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("major")
		e.Int(s.Major)
	}
	{
		e.FieldStart("minor")
		e.Int(s.Minor)
	}
	{
		e.FieldStart("patch")
		e.Int(s.Patch)
	}
	{
		if s.Suffix.Set {
			e.FieldStart("suffix")
			s.Suffix.Encode(e)
		}
	}
}

var jsonFieldsNameOfVersion = [4]string{
	0: "major",
	1: "minor",
	2: "patch",
	3: "suffix",
}

// Decode decodes Version from json.
func (s *Version) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Version to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "major":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int()
				s.Major = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"major\"")
			}
		case "minor":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Int()
				s.Minor = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"minor\"")
			}
		case "patch":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Int()
				s.Patch = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"patch\"")
			}
		case "suffix":
			if err := func() error {
				s.Suffix.Reset()
				if err := s.Suffix.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"suffix\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Version")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfVersion) {
					name = jsonFieldsNameOfVersion[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Version) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Version) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
