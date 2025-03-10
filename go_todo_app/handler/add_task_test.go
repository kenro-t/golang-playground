package handler

import (
	"bytes"
	"go_todo_app/entity"
	"go_todo_app/store"
	"go_todo_app/testutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestAddTask(t *testing.T) {
	type want struct {
		status  int
		rspFile string
	}
	tests := map[string]struct {
		reqFile string
		want    want
	}{
		"ok": {
			reqFile: "testdata/add_task/ok_req.json",
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/add_task/ok_rsp.json",
			},
		},
		"badRequest": {
			reqFile: "testdata/add_task/bad_req.json",
			want: want{
				status:  http.StatusBadRequest,
				rspFile: "testdata/add_task/bad_rsp.json",
			},
		},
	}
	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodPost,
				"/tasks",
				bytes.NewReader(testutil.LoadFile(t, tt.reqFile)),
			)

			sut := AddTask{Store: &store.TaskStore{
				Tasks: map[entity.TaskID]*entity.Task{},
			}, Validator: validator.New()}
			sut.ServeHTTP(w, r)

			resp := w.Result()
			testutil.AssertResponse(t,
				resp, tt.want.status, testutil.LoadFile(t, tt.want.rspFile),
			)
		})
	}
}
