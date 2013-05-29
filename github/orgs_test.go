// Copyright 2013 Google. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file or at
// https://developers.google.com/open-source/licenses/bsd

package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestOrganizationsService_List_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/user/orgs", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Request method = %v, want %v", r.Method, "GET")
		}
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	orgs, err := client.Organizations.List("", nil)
	if err != nil {
		t.Errorf("Organizations.List returned error: %v", err)
	}

	want := []Organization{Organization{ID: 1}, Organization{ID: 2}}
	if !reflect.DeepEqual(orgs, want) {
		t.Errorf("Organizations.List returned %+v, want %+v", orgs, want)
	}
}

func TestOrganizationsService_List_specifiedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/u/orgs", func(w http.ResponseWriter, r *http.Request) {
		var v string
		if r.Method != "GET" {
			t.Errorf("Request method = %v, want %v", r.Method, "GET")
		}
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
		if v = r.FormValue("page"); v != "2" {
			t.Errorf("Request page parameter = %v, want %v", v, "2")
		}
	})

	opt := &ListOptions{2}
	orgs, err := client.Organizations.List("u", opt)
	if err != nil {
		t.Errorf("Organizations.List returned error: %v", err)
	}

	want := []Organization{Organization{ID: 1}, Organization{ID: 2}}
	if !reflect.DeepEqual(orgs, want) {
		t.Errorf("Organizations.List returned %+v, want %+v", orgs, want)
	}
}

func TestOrganizationsService_List_invalidUser(t *testing.T) {
	_, err := client.Organizations.List("%", nil)
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

func TestOrganizationsService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/orgs/o", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Request method = %v, want %v", r.Method, "GET")
		}
		fmt.Fprint(w, `{"id":1, "login":"l", "url":"u", "avatar_url": "a", "location":"l"}`)
	})

	org, err := client.Organizations.Get("o")
	if err != nil {
		t.Errorf("Organizations.Get returned error: %v", err)
	}

	want := &Organization{ID: 1, Login: "l", URL: "u", AvatarURL: "a", Location: "l"}
	if !reflect.DeepEqual(org, want) {
		t.Errorf("Organizations.Get returned %+v, want %+v", org, want)
	}
}

func TestOrganizationsService_Get_invalidOrg(t *testing.T) {
	_, err := client.Organizations.Get("%")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

func TestOrganizationsService_Edit(t *testing.T) {
	setup()
	defer teardown()

	input := &Organization{Login: "l"}

	mux.HandleFunc("/orgs/o", func(w http.ResponseWriter, r *http.Request) {
		v := new(Organization)
		json.NewDecoder(r.Body).Decode(v)

		if r.Method != "PATCH" {
			t.Errorf("Request method = %v, want %v", r.Method, "GET")
		}
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"id":1}`)
	})

	org, err := client.Organizations.Edit("o", input)
	if err != nil {
		t.Errorf("Organizations.Edit returned error: %v", err)
	}

	want := &Organization{ID: 1}
	if !reflect.DeepEqual(org, want) {
		t.Errorf("Organizations.Edit returned %+v, want %+v", org, want)
	}
}

func TestOrganizationsService_Edit_invalidOrg(t *testing.T) {
	_, err := client.Organizations.Edit("%", nil)
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

func TestOrganizationsService_ListMembers(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/orgs/o/members", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Request method = %v, want %v", r.Method, "GET")
		}
		fmt.Fprint(w, `[{"id":1}]`)
	})

	members, err := client.Organizations.ListMembers("o")
	if err != nil {
		t.Errorf("Organizations.ListMembers returned error: %v", err)
	}

	want := []User{User{ID: 1}}
	if !reflect.DeepEqual(members, want) {
		t.Errorf("Organizations.ListMembers returned %+v, want %+v", members, want)
	}
}

func TestOrganizationsService_ListMembers_invalidOrg(t *testing.T) {
	_, err := client.Organizations.ListMembers("%")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

func TestOrganizationsService_ListPublicMembers(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/orgs/o/public_members", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Request method = %v, want %v", r.Method, "GET")
		}
		fmt.Fprint(w, `[{"id":1}]`)
	})

	members, err := client.Organizations.ListPublicMembers("o")
	if err != nil {
		t.Errorf("Organizations.ListPublicMembers returned error: %v", err)
	}

	want := []User{User{ID: 1}}
	if !reflect.DeepEqual(members, want) {
		t.Errorf("Organizations.ListPublicMembers returned %+v, want %+v", members, want)
	}
}

func TestOrganizationsService_ListPublicMembers_invalidOrg(t *testing.T) {
	_, err := client.Organizations.ListPublicMembers("%")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

func TestOrganizationsService_CheckMembership(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/orgs/o/members/u", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Request method = %v, want %v", r.Method, "GET")
		}
	})

	member, err := client.Organizations.CheckMembership("o", "u")
	if err != nil {
		t.Errorf("Organizations.CheckMembership returned error: %v", err)
	}
	want := true
	if member != want {
		t.Errorf("Organizations.CheckMembership returned %+v, want %+v", member, want)
	}
}

// ensure that a 404 response is interpreted as "false" and not an error
func TestOrganizationsService_CheckMembership_notMember(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/orgs/o/members/u", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Request method = %v, want %v", r.Method, "GET")
		}
		w.WriteHeader(http.StatusNotFound)
	})

	member, err := client.Organizations.CheckMembership("o", "u")
	if err != nil {
		t.Errorf("Organizations.CheckMembership returned error: %+v", err)
	}
	want := false
	if member != want {
		t.Errorf("Organizations.CheckMembership returned %+v, want %+v", member, want)
	}
}

// ensure that a 400 response is interpreted as an actual error, and not simply
// as "false" like the above case of a 404
func TestOrganizationsService_CheckMembership_error(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/orgs/o/members/u", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Request method = %v, want %v", r.Method, "GET")
		}
		http.Error(w, "BadRequest", http.StatusBadRequest)
	})

	member, err := client.Organizations.CheckMembership("o", "u")
	if err == nil {
		t.Errorf("Expected HTTP 400 response")
	}
	want := false
	if member != want {
		t.Errorf("Organizations.CheckMembership returned %+v, want %+v", member, want)
	}
}

func TestOrganizationsService_CheckMembership_invalidOrg(t *testing.T) {
	_, err := client.Organizations.CheckMembership("%", "u")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

func TestOrganizationsService_CheckPublicMembership(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/orgs/o/public_members/u", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Request method = %v, want %v", r.Method, "GET")
		}
	})

	member, err := client.Organizations.CheckPublicMembership("o", "u")
	if err != nil {
		t.Errorf("Organizations.CheckPublicMembership returned error: %v", err)
	}
	want := true
	if member != want {
		t.Errorf("Organizations.CheckPublicMembership returned %+v, want %+v", member, want)
	}
}

// ensure that a 404 response is interpreted as "false" and not an error
func TestOrganizationsService_CheckPublicMembership_notMember(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/orgs/o/public_members/u", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Request method = %v, want %v", r.Method, "GET")
		}
		w.WriteHeader(http.StatusNotFound)
	})

	member, err := client.Organizations.CheckPublicMembership("o", "u")
	if err != nil {
		t.Errorf("Organizations.CheckPublicMembership returned error: %v", err)
	}
	want := false
	if member != want {
		t.Errorf("Organizations.CheckPublicMembership returned %+v, want %+v", member, want)
	}
}

// ensure that a 400 response is interpreted as an actual error, and not simply
// as "false" like the above case of a 404
func TestOrganizationsService_CheckPublicMembership_error(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/orgs/o/public_members/u", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Request method = %v, want %v", r.Method, "GET")
		}
		http.Error(w, "BadRequest", http.StatusBadRequest)
	})

	member, err := client.Organizations.CheckPublicMembership("o", "u")
	if err == nil {
		t.Errorf("Expected HTTP 400 response")
	}
	want := false
	if member != want {
		t.Errorf("Organizations.CheckPublicMembership returned %+v, want %+v", member, want)
	}
}

func TestOrganizationsService_CheckPublicMembership_invalidOrg(t *testing.T) {
	_, err := client.Organizations.CheckPublicMembership("%", "u")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

func TestOrganizationsService_RemoveMember(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/orgs/o/members/u", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("Request method = %v, want %v", r.Method, "DELETE")
		}
	})

	err := client.Organizations.RemoveMember("o", "u")
	if err != nil {
		t.Errorf("Organizations.RemoveMember returned error: %v", err)
	}
}

func TestOrganizationsService_RemoveMember_invalidOrg(t *testing.T) {
	err := client.Organizations.RemoveMember("%", "u")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

func TestOrganizationsService_ListTeams(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/orgs/o/teams", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Request method = %v, want %v", r.Method, "GET")
		}
		fmt.Fprint(w, `[{"id":1}]`)
	})

	teams, err := client.Organizations.ListTeams("o")
	if err != nil {
		t.Errorf("Organizations.ListTeams returned error: %v", err)
	}

	want := []Team{Team{ID: 1}}
	if !reflect.DeepEqual(teams, want) {
		t.Errorf("Organizations.ListTeams returned %+v, want %+v", teams, want)
	}
}

func TestOrganizationsService_ListTeams_invalidOrg(t *testing.T) {
	_, err := client.Organizations.ListTeams("%")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

func TestOrganizationsService_AddTeamMember(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/teams/1/members/u", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("Request method = %v, want %v", r.Method, "GET")
		}
	})

	err := client.Organizations.AddTeamMember(1, "u")
	if err != nil {
		t.Errorf("Organizations.AddTeamMember returned error: %v", err)
	}
}

func TestOrganizationsService_AddTeamMember_invalidUser(t *testing.T) {
	err := client.Organizations.AddTeamMember(1, "%")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

func TestOrganizationsService_RemoveTeamMember(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/teams/1/members/u", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("Request method = %v, want %v", r.Method, "GET")
		}
	})

	err := client.Organizations.RemoveTeamMember(1, "u")
	if err != nil {
		t.Errorf("Organizations.RemoveTeamMember returned error: %v", err)
	}
}

func TestOrganizationsService_RemoveTeamMember_invalidUser(t *testing.T) {
	err := client.Organizations.RemoveTeamMember(1, "%")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

func TestOrganizationsService_PublicizeMembership(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/orgs/o/public_members/u", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("Request method = %v, want %v", r.Method, "GET")
		}
	})

	err := client.Organizations.PublicizeMembership("o", "u")
	if err != nil {
		t.Errorf("Organizations.PublicizeMembership returned error: %v", err)
	}
}

func TestOrganizationsService_PublicizeMembership_invalidOrg(t *testing.T) {
	err := client.Organizations.PublicizeMembership("%", "u")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

func TestOrganizationsService_ConcealMembership(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/orgs/o/public_members/u", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("Request method = %v, want %v", r.Method, "GET")
		}
	})

	err := client.Organizations.ConcealMembership("o", "u")
	if err != nil {
		t.Errorf("Organizations.ConcealMembership returned error: %v", err)
	}
}

func TestOrganizationsService_ConcealMembership_invalidOrg(t *testing.T) {
	err := client.Organizations.ConcealMembership("%", "u")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}