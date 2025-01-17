// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func encodeFilesGetResponse(response []string, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	e := jx.GetEncoder()
	e.ArrStart()
	for _, elem := range response {
		e.Str(elem)
	}
	e.ArrEnd()
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}
	return nil
}

func encodeGetFilesRootRelpathResponse(response GetFilesRootRelpathRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *Directory:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *GetFilesRootRelpathNotFound:
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeGetPlayerPauseResponse(response *Pause, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	e := jx.GetEncoder()
	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}
	return nil
}

func encodeGetPlayerPositionResponse(response PlaybackPosition, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	span.SetStatus(codes.Ok, http.StatusText(200))

	e := jx.GetEncoder()
	response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}
	return nil
}

func encodeGetStatusResponse(response GetStatusRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *GetStatusOK:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *GetStatusServiceUnavailable:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(503)
		span.SetStatus(codes.Error, http.StatusText(503))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodePlayerStartPostResponse(response PlayerStartPostRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PlayerStartPostAccepted:
		w.WriteHeader(202)
		span.SetStatus(codes.Ok, http.StatusText(202))

		return nil

	case *PlayerStartPostBadRequest:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	case *PlayerStartPostNotFound:
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodePutPlayerPauseResponse(response *PutPlayerPauseAccepted, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(202)
	span.SetStatus(codes.Ok, http.StatusText(202))

	return nil
}

func encodePutPlayerPositionResponse(response *PutPlayerPositionAccepted, w http.ResponseWriter, span trace.Span) error {
	w.WriteHeader(202)
	span.SetStatus(codes.Ok, http.StatusText(202))

	return nil
}
