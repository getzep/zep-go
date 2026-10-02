// This test table comes from spec 3 section 4.2 of the v4 public API specification.
// This file is hand-written and listed in .fernignore.

package zep_test

import (
	"context"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"testing"

	zep "github.com/getzep/zep-go/v4"
	sdkclient "github.com/getzep/zep-go/v4/client"
	"github.com/getzep/zep-go/v4/core"
	"github.com/getzep/zep-go/v4/option"
	"github.com/google/uuid"
)

type section42Operation struct {
	name      string
	method    string
	path      string
	paginated bool
	postRead  bool
}

// Source: spec 3 section 4.2 (endpoint map). Fields: SDK method, HTTP method,
// path, paginated (P), POST read (spec 3 section 2.9).
var section42Operations = []section42Operation{
	{"project.get", "GET", "/project", false, false},
	{"project.update", "PATCH", "/project", false, false},
	{"project.get_content_policy", "GET", "/project/content-policy", false, false},
	{"project.set_content_policy", "PUT", "/project/content-policy", false, false},
	{"project.list_content_policy_revisions", "GET", "/project/content-policy/revisions", true, false},
	{"project.get_content_policy_revision", "GET", "/project/content-policy/revisions/{revision_uuid}", false, false},
	{"project.get_instructions", "GET", "/project/instructions", false, false},
	{"project.set_instructions", "PUT", "/project/instructions", false, false},
	{"project.get_observation_steering", "GET", "/project/observation-steering", false, false},
	{"project.set_observation_steering", "PUT", "/project/observation-steering", false, false},
	{"project.get_ontology", "GET", "/project/ontology", false, false},
	{"project.set_ontology", "PUT", "/project/ontology", false, false},
	{"project.get_user_summary_instructions", "GET", "/project/user-summary-instructions", false, false},
	{"project.set_user_summary_instructions", "PUT", "/project/user-summary-instructions", false, false},
	{"context.create_template", "POST", "/context-templates", false, false},
	{"context.list_templates", "POST", "/context-templates/list", true, true},
	{"context.delete_template", "DELETE", "/context-templates/{template_uuid}", false, false},
	{"context.get_template", "GET", "/context-templates/{template_uuid}", false, false},
	{"context.update_template", "PUT", "/context-templates/{template_uuid}", false, false},
	{"agent.create", "POST", "/agents", false, false},
	{"agent.list", "POST", "/agents/list", true, true},
	{"agent.delete", "DELETE", "/agents/{agent_uuid}", false, false},
	{"agent.get", "GET", "/agents/{agent_uuid}", false, false},
	{"agent.update", "PATCH", "/agents/{agent_uuid}", false, false},
	{"agent.declare_breaking_change", "POST", "/agents/{agent_uuid}/breaking-changes", false, false},
	{"agent.get_context", "POST", "/agents/{agent_uuid}/context", false, true},
	{"agent.split.plan", "POST", "/agents/{agent_uuid}/split-plan", false, false},
	{"agent.literal_policy.get", "GET", "/agents/{agent_uuid}/literal-policy", false, false},
	{"agent.literal_policy.update", "PUT", "/agents/{agent_uuid}/literal-policy", false, false},
	{"agent.skill.candidate.list", "GET", "/agents/{agent_uuid}/skill-candidates", true, false},
	{"agent.skill.candidate.get", "GET", "/agents/{agent_uuid}/skill-candidates/{review_uuid}", false, false},
	{"agent.skill.evaluation.create_for_candidate", "POST", "/agents/{agent_uuid}/skill-candidates/{review_uuid}/candidates/{candidate_uuid}/evaluations", false, false},
	{"agent.learning.get", "GET", "/agents/{agent_uuid}/learning", false, false},
	{"agent.learning.list_runs", "GET", "/agents/{agent_uuid}/learning-runs", true, false},
	{"agent.skill.create", "POST", "/agents/{agent_uuid}/skills", false, false},
	{"agent.skill.import_package", "POST", "/agents/{agent_uuid}/skills/import", false, false},
	{"agent.skill.list", "POST", "/agents/{agent_uuid}/skills/list", true, true},
	{"agent.skill.search", "POST", "/agents/{agent_uuid}/skills/search", true, true},
	{"agent.skill.get", "GET", "/agents/{agent_uuid}/skills/{skill_uuid}", false, false},
	{"agent.skill.publication.get", "GET", "/agents/{agent_uuid}/skills/{skill_uuid}/publication", false, false},
	{"agent.skill.publication.lookup", "GET", "/agents/{agent_uuid}/skill-publications", false, false},
	{"agent.skill.use.create", "POST", "/agents/{agent_uuid}/skills/{skill_uuid}/uses", false, false},
	{"agent.skill.use.add_outcome", "POST", "/agents/{agent_uuid}/skills/{skill_uuid}/uses/{use_uuid}/outcome", false, false},
	{"agent.skill.approve", "POST", "/agents/{agent_uuid}/skills/{skill_uuid}/approve", false, false},
	{"agent.skill.evaluation.create", "POST", "/agents/{agent_uuid}/skills/{skill_uuid}/evaluations", false, false},
	{"agent.skill.evidence.list", "GET", "/agents/{agent_uuid}/skills/{skill_uuid}/evidence", true, false},
	{"agent.skill.relation.list", "GET", "/agents/{agent_uuid}/skills/{skill_uuid}/relations", true, false},
	{"agent.skill.version.restore_version", "POST", "/agents/{agent_uuid}/skills/{skill_uuid}/restore-version", false, false},
	{"agent.skill.retire", "POST", "/agents/{agent_uuid}/skills/{skill_uuid}/retire", false, false},
	{"agent.skill.version.list", "GET", "/agents/{agent_uuid}/skills/{skill_uuid}/versions", true, false},
	{"agent.skill.create_version", "POST", "/agents/{agent_uuid}/skills/{skill_uuid}/versions", false, false},
	{"agent.skill.version.compare", "POST", "/agents/{agent_uuid}/skills/{skill_uuid}/versions/compare", false, true},
	{"agent.skill.version.get", "GET", "/agents/{agent_uuid}/skills/{skill_uuid}/versions/{version}", false, false},
	{"agent.skill.export.create", "POST", "/agents/{agent_uuid}/skills/{skill_uuid}/versions/{version}/export", false, false},
	{"agent.skill.export.get", "GET", "/agents/{agent_uuid}/skills/{skill_uuid}/versions/{version}/export/{task_uuid}", false, false},
	{"agent.trajectory.create", "POST", "/agents/{agent_uuid}/trajectories", false, false},
	{"agent.trajectory.list", "POST", "/agents/{agent_uuid}/trajectories/list", true, true},
	{"agent.trajectory.get", "GET", "/agents/{agent_uuid}/trajectories/{trajectory_uuid}", false, false},
	{"agent.trajectory.update", "PATCH", "/agents/{agent_uuid}/trajectories/{trajectory_uuid}", false, false},
	{"agent.trajectory.delete", "DELETE", "/agents/{agent_uuid}/trajectories/{trajectory_uuid}", false, false},
	{"agent.trajectory.abandon", "POST", "/agents/{agent_uuid}/trajectories/{trajectory_uuid}/abandon", false, false},
	{"agent.trajectory.close", "POST", "/agents/{agent_uuid}/trajectories/{trajectory_uuid}/close", false, false},
	{"agent.trajectory.correct_task_family", "POST", "/agents/{agent_uuid}/trajectories/{trajectory_uuid}/correct-task-family", false, false},
	{"agent.trajectory.list_events", "GET", "/agents/{agent_uuid}/trajectories/{trajectory_uuid}/events", true, false},
	{"agent.trajectory.append_event", "POST", "/agents/{agent_uuid}/trajectories/{trajectory_uuid}/events", false, false},
	{"agent.trajectory.delete_event", "DELETE", "/agents/{agent_uuid}/trajectories/{trajectory_uuid}/events/{event_uuid}", false, false},
	{"agent.trajectory.reopen", "POST", "/agents/{agent_uuid}/trajectories/{trajectory_uuid}/reopen", false, false},
	{"agent.trajectory.get_summary", "GET", "/agents/{agent_uuid}/trajectories/{trajectory_uuid}/summary", false, false},
	{"agent.trajectory.list_summary_versions", "GET", "/agents/{agent_uuid}/trajectories/{trajectory_uuid}/summary/versions", true, false},
	{"agent.verifier.list", "POST", "/agents/{agent_uuid}/verifiers/list", true, true},
	{"agent.verifier.get", "GET", "/agents/{agent_uuid}/verifiers/{verifier_uuid}", false, false},
	{"agent.verifier.update", "PATCH", "/agents/{agent_uuid}/verifiers/{verifier_uuid}", false, false},
	{"agent.verifier.invalidate_evidence", "POST", "/agents/{agent_uuid}/verifiers/{verifier_uuid}/invalidate-evidence", false, false},
	{"agent.verifier.revoke", "POST", "/agents/{agent_uuid}/verifiers/{verifier_uuid}/revoke", false, false},
	{"user.create", "POST", "/users", false, false},
	{"user.list", "POST", "/users/list", true, true},
	{"user.lookup", "POST", "/users/lookup", false, true},
	{"user.delete", "DELETE", "/users/{user_uuid}", false, false},
	{"user.get", "GET", "/users/{user_uuid}", false, false},
	{"user.update", "PATCH", "/users/{user_uuid}", false, false},
	{"user.get_node", "GET", "/users/{user_uuid}/node", false, false},
	{"user.get_summary_instructions", "GET", "/users/{user_uuid}/summary-instructions", false, false},
	{"user.set_summary_instructions", "PUT", "/users/{user_uuid}/summary-instructions", false, false},
	{"user_group.list_for_user", "GET", "/users/{user_uuid}/user-groups", true, false},
	{"thread.list", "GET", "/threads", true, false},
	{"thread.create", "POST", "/threads", false, false},
	{"thread.lookup", "POST", "/threads/lookup", false, true},
	{"thread.delete", "DELETE", "/threads/{thread_uuid}", false, false},
	{"thread.get", "GET", "/threads/{thread_uuid}", false, false},
	{"thread.get_context", "GET", "/threads/{thread_uuid}/context", false, false},
	{"thread.list_episodes", "GET", "/threads/{thread_uuid}/episodes", true, false},
	{"thread.list_messages", "GET", "/threads/{thread_uuid}/messages", true, false},
	{"thread.add_messages", "POST", "/threads/{thread_uuid}/messages", false, false},
	{"thread.message.get", "GET", "/threads/{thread_uuid}/messages/{message_uuid}", false, false},
	{"thread.message.update", "PATCH", "/threads/{thread_uuid}/messages/{message_uuid}", false, false},
	{"thread.get_summary", "GET", "/threads/{thread_uuid}/summary", false, false},
	{"graph.create", "POST", "/graphs", false, false},
	{"graph.list", "POST", "/graphs/list", true, true},
	{"graph.lookup", "POST", "/graphs/lookup", false, true},
	{"graph.delete", "DELETE", "/graphs/{graph_uuid}", false, false},
	{"graph.get", "GET", "/graphs/{graph_uuid}", false, false},
	{"graph.update", "PATCH", "/graphs/{graph_uuid}", false, false},
	{"graph.clone", "POST", "/graphs/{graph_uuid}/clone", false, false},
	{"graph.get_context", "POST", "/graphs/{graph_uuid}/context", false, true},
	{"graph.get_content_policy", "GET", "/graphs/{graph_uuid}/content-policy", false, false},
	{"graph.list_content_policy_events", "POST", "/graphs/{graph_uuid}/content-policy/events/list", true, true},
	{"graph.document_summary.list", "POST", "/graphs/{graph_uuid}/document-summaries/list", true, true},
	{"graph.episode.list_for_document", "GET", "/graphs/{graph_uuid}/documents/{document_id}/episodes", true, false},
	{"graph.edge.add", "POST", "/graphs/{graph_uuid}/edges", false, false},
	{"graph.edge.list", "POST", "/graphs/{graph_uuid}/edges/list", true, true},
	{"graph.edge.delete", "DELETE", "/graphs/{graph_uuid}/edges/{edge_uuid}", false, false},
	{"graph.edge.get", "GET", "/graphs/{graph_uuid}/edges/{edge_uuid}", false, false},
	{"graph.edge.update", "PATCH", "/graphs/{graph_uuid}/edges/{edge_uuid}", false, false},
	{"graph.episode.add", "POST", "/graphs/{graph_uuid}/episodes", false, false},
	{"graph.episode.list", "POST", "/graphs/{graph_uuid}/episodes/list", true, true},
	{"graph.episode.delete", "DELETE", "/graphs/{graph_uuid}/episodes/{episode_uuid}", false, false},
	{"graph.episode.get", "GET", "/graphs/{graph_uuid}/episodes/{episode_uuid}", false, false},
	{"graph.episode.update", "PATCH", "/graphs/{graph_uuid}/episodes/{episode_uuid}", false, false},
	{"graph.hyperedge.add", "POST", "/graphs/{graph_uuid}/hyperedges", false, false},
	{"graph.hyperedge.list", "POST", "/graphs/{graph_uuid}/hyperedges/list", true, true},
	{"graph.hyperedge.delete", "DELETE", "/graphs/{graph_uuid}/hyperedges/{hyperedge_uuid}", false, false},
	{"graph.hyperedge.get", "GET", "/graphs/{graph_uuid}/hyperedges/{hyperedge_uuid}", false, false},
	{"graph.hyperedge.update", "PATCH", "/graphs/{graph_uuid}/hyperedges/{hyperedge_uuid}", false, false},
	{"graph.hyperedge.create_edge", "POST", "/graphs/{graph_uuid}/hyperedges/{hyperedge_uuid}/edges", false, false},
	{"graph.hyperedge.delete_edge", "DELETE", "/graphs/{graph_uuid}/hyperedges/{hyperedge_uuid}/edges/{edge_uuid}", false, false},
	{"graph.get_instructions", "GET", "/graphs/{graph_uuid}/instructions", false, false},
	{"graph.set_instructions", "PUT", "/graphs/{graph_uuid}/instructions", false, false},
	{"graph.node.add", "POST", "/graphs/{graph_uuid}/nodes", false, false},
	{"graph.node.list", "POST", "/graphs/{graph_uuid}/nodes/list", true, true},
	{"graph.node.delete", "DELETE", "/graphs/{graph_uuid}/nodes/{node_uuid}", false, false},
	{"graph.node.get", "GET", "/graphs/{graph_uuid}/nodes/{node_uuid}", false, false},
	{"graph.node.update", "PATCH", "/graphs/{graph_uuid}/nodes/{node_uuid}", false, false},
	{"graph.node.list_neighbors", "POST", "/graphs/{graph_uuid}/nodes/{node_uuid}/neighbors", true, true},
	{"graph.get_observation_steering", "GET", "/graphs/{graph_uuid}/observation-steering", false, false},
	{"graph.set_observation_steering", "PUT", "/graphs/{graph_uuid}/observation-steering", false, false},
	{"graph.observation.list", "POST", "/graphs/{graph_uuid}/observations/list", true, true},
	{"graph.observation.get", "GET", "/graphs/{graph_uuid}/observations/{observation_uuid}", false, false},
	{"graph.get_ontology", "GET", "/graphs/{graph_uuid}/ontology", false, false},
	{"graph.set_ontology", "PUT", "/graphs/{graph_uuid}/ontology", false, false},
	{"graph.search_edges", "POST", "/graphs/{graph_uuid}/search/edges", true, true},
	{"graph.search_episodes", "POST", "/graphs/{graph_uuid}/search/episodes", true, true},
	{"graph.search_nodes", "POST", "/graphs/{graph_uuid}/search/nodes", true, true},
	{"graph.search_observations", "POST", "/graphs/{graph_uuid}/search/observations", true, true},
	{"graph.search_thread_summaries", "POST", "/graphs/{graph_uuid}/search/thread-summaries", true, true},
	{"graph.get_subgraph", "POST", "/graphs/{graph_uuid}/subgraph", false, true},
	{"graph.thread_summary.list", "POST", "/graphs/{graph_uuid}/thread-summaries/list", true, true},
	{"graph.warm", "POST", "/graphs/{graph_uuid}/warm", false, false},
	{"batch.list", "GET", "/batches", true, false},
	{"batch.create", "POST", "/batches", false, false},
	{"batch.delete", "DELETE", "/batches/{batch_uuid}", false, false},
	{"batch.get", "GET", "/batches/{batch_uuid}", false, false},
	{"batch.list_items", "GET", "/batches/{batch_uuid}/items", true, false},
	{"batch.add_items", "POST", "/batches/{batch_uuid}/items", false, false},
	{"batch.process", "POST", "/batches/{batch_uuid}/process", false, false},
	{"task.list", "GET", "/tasks", true, false},
	{"task.get", "GET", "/tasks/{task_uuid}", false, false},
	{"user_group.create", "POST", "/user-groups", false, false},
	{"user_group.list", "POST", "/user-groups/list", true, true},
	{"user_group.delete", "DELETE", "/user-groups/{group_uuid}", false, false},
	{"user_group.get", "GET", "/user-groups/{group_uuid}", false, false},
	{"user_group.update", "PATCH", "/user-groups/{group_uuid}", false, false},
	{"user_group.list_member_candidates", "POST", "/user-groups/{group_uuid}/member-candidates/list", true, true},
	{"user_group.add_members", "POST", "/user-groups/{group_uuid}/members", false, false},
	{"user_group.list_members", "POST", "/user-groups/{group_uuid}/members/list", true, true},
	{"user_group.remove_members", "POST", "/user-groups/{group_uuid}/members/remove", false, false},
	{"user_group.remove_member", "DELETE", "/user-groups/{group_uuid}/members/{user_uuid}", false, false},
	{"lookup.batch", "POST", "/lookup", false, true},
}

var missingFromAlpha5 = map[string]bool{
	"project.get_content_policy":                  true,
	"project.set_content_policy":                  true,
	"project.list_content_policy_revisions":       true,
	"project.get_content_policy_revision":         true,
	"agent.create":                                true,
	"agent.list":                                  true,
	"agent.delete":                                true,
	"agent.get":                                   true,
	"agent.update":                                true,
	"agent.declare_breaking_change":               true,
	"agent.get_context":                           true,
	"agent.split.plan":                            true,
	"agent.literal_policy.get":                    true,
	"agent.literal_policy.update":                 true,
	"agent.skill.candidate.list":                  true,
	"agent.skill.candidate.get":                   true,
	"agent.skill.evaluation.create_for_candidate": true,
	"agent.learning.get":                          true,
	"agent.learning.list_runs":                    true,
	"agent.skill.create":                          true,
	"agent.skill.import_package":                  true,
	"agent.skill.list":                            true,
	"agent.skill.search":                          true,
	"agent.skill.get":                             true,
	"agent.skill.publication.get":                 true,
	"agent.skill.publication.lookup":              true,
	"agent.skill.use.create":                      true,
	"agent.skill.use.add_outcome":                 true,
	"agent.skill.approve":                         true,
	"agent.skill.evaluation.create":               true,
	"agent.skill.evidence.list":                   true,
	"agent.skill.relation.list":                   true,
	"agent.skill.version.restore_version":         true,
	"agent.skill.retire":                          true,
	"agent.skill.version.list":                    true,
	"agent.skill.create_version":                  true,
	"agent.skill.version.compare":                 true,
	"agent.skill.version.get":                     true,
	"agent.skill.export.create":                   true,
	"agent.skill.export.get":                      true,
	"agent.trajectory.create":                     true,
	"agent.trajectory.list":                       true,
	"agent.trajectory.get":                        true,
	"agent.trajectory.update":                     true,
	"agent.trajectory.delete":                     true,
	"agent.trajectory.abandon":                    true,
	"agent.trajectory.close":                      true,
	"agent.trajectory.correct_task_family":        true,
	"agent.trajectory.list_events":                true,
	"agent.trajectory.append_event":               true,
	"agent.trajectory.delete_event":               true,
	"agent.trajectory.reopen":                     true,
	"agent.trajectory.get_summary":                true,
	"agent.trajectory.list_summary_versions":      true,
	"agent.verifier.list":                         true,
	"agent.verifier.get":                          true,
	"agent.verifier.update":                       true,
	"agent.verifier.invalidate_evidence":          true,
	"agent.verifier.revoke":                       true,
	"graph.get_content_policy":                    true,
	"graph.list_content_policy_events":            true,
	"graph.hyperedge.add":                         true,
	"graph.hyperedge.list":                        true,
	"graph.hyperedge.delete":                      true,
	"graph.hyperedge.get":                         true,
	"graph.hyperedge.update":                      true,
	"graph.hyperedge.create_edge":                 true,
	"graph.hyperedge.delete_edge":                 true,
}

var alpha5PostReadExposesIdempotency = map[string]bool{
	"context.list_templates":            true,
	"user.list":                         true,
	"user.lookup":                       true,
	"thread.lookup":                     true,
	"graph.list":                        true,
	"graph.lookup":                      true,
	"graph.get_context":                 true,
	"graph.document_summary.list":       true,
	"graph.edge.list":                   true,
	"graph.episode.list":                true,
	"graph.node.list":                   true,
	"graph.node.list_neighbors":         true,
	"graph.observation.list":            true,
	"graph.search_edges":                true,
	"graph.search_episodes":             true,
	"graph.search_nodes":                true,
	"graph.search_observations":         true,
	"graph.search_thread_summaries":     true,
	"graph.get_subgraph":                true,
	"graph.thread_summary.list":         true,
	"user_group.list":                   true,
	"user_group.list_member_candidates": true,
	"user_group.list_members":           true,
	"lookup.batch":                      true,
}

var excludedFromSDK = map[string]bool{
	"user_group.list_policy_sets":    true, // docs audience
	"user_group.attach_policy_set":   true, // docs audience
	"user_group.detach_policy_set":   true, // docs audience
	"abac.list_api_keys":             true, // /abac administrative plane (zepctl only)
	"abac.list_api_key_policy_sets":  true, // /abac administrative plane (zepctl only)
	"abac.attach_api_key_policy_set": true, // /abac administrative plane (zepctl only)
	"abac.detach_api_key_policy_set": true, // /abac administrative plane (zepctl only)
	"abac.get_api_key_settings":      true, // /abac administrative plane (zepctl only)
	"abac.set_api_key_settings":      true, // /abac administrative plane (zepctl only)
	"abac.evaluate_policy":           true, // /abac administrative plane (zepctl only)
	"abac.explain_policy":            true, // /abac administrative plane (zepctl only)
	"abac.list_policy_sets":          true, // /abac administrative plane (zepctl only)
	"abac.create_policy_set":         true, // /abac administrative plane (zepctl only)
	"abac.validate_policy_set":       true, // /abac administrative plane (zepctl only)
	"abac.delete_policy_set":         true, // /abac administrative plane (zepctl only)
	"abac.get_policy_set":            true, // /abac administrative plane (zepctl only)
	"abac.update_policy_set":         true, // /abac administrative plane (zepctl only)
	"abac.detach_retention_target":   true, // /abac administrative plane (zepctl only)
	"abac.list_retention_targets":    true, // /abac administrative plane (zepctl only)
	"abac.attach_retention_target":   true, // /abac administrative plane (zepctl only)
}

const (
	d1Reason        = "The generator configuration does not enable automatic Idempotency-Key generation (spec 3 section 14.6), so a state-changing call without a caller key sends no Idempotency-Key."
	d2Reason        = "The generated IdempotentRequestOptions.ToHeader sends the key with a '*' prefix, so a caller key is not sent unchanged."
	callerKey       = "contract-caller-key"
	projectUUID     = "00000000-0000-4000-8000-000000000001"
	recencyBiasType = "V4GraphContextRequestRecencyBias"
)

var (
	lastUpperBoundary  = regexp.MustCompile(`([A-Z]+)([A-Z][a-z])`)
	lowerUpperBoundary = regexp.MustCompile(`([a-z0-9])([A-Z])`)
	nonNameCharacters  = regexp.MustCompile(`[^a-zA-Z0-9]+`)
)

func sdkVersion() string {
	return core.NewRequestOptions().ToHeader().Get("X-Fern-SDK-Version")
}

func missingOperationReason(operation string) string {
	if sdkVersion() == "v4.0.0-alpha.5" && missingFromAlpha5[operation] {
		return operation + ": absent from the 4.0.0-alpha.5 generated code; present in the current spec 3 contract"
	}
	return ""
}

func contractGapReason(operation string, postRead bool) string {
	if reason := missingOperationReason(operation); reason != "" {
		return reason
	}
	if sdkVersion() == "v4.0.0-alpha.5" && postRead && alpha5PostReadExposesIdempotency[operation] {
		return operation + ": POST read; the 4.0.0-alpha.5 contract marks it idempotent"
	}
	return ""
}

func expectContractFailure(t *testing.T, reason string, err error) {
	t.Helper()
	if err == nil {
		t.Fatalf("%s is fixed; remove the entry", reason)
	}
	t.Logf("expected failure: %s: %v", reason, err)
}

func newContractClient(server *httptest.Server, maxAttempts uint, options ...option.RequestOption) *sdkclient.Client {
	requestOptions := []option.RequestOption{option.WithMaxAttempts(maxAttempts)}
	if server != nil {
		requestOptions = append(requestOptions, option.WithBaseURL(server.URL))
	}
	requestOptions = append(requestOptions, options...)
	return sdkclient.NewClient(requestOptions...)
}

func goOperationName(operation string) string {
	segments := strings.Split(operation, ".")
	for index, segment := range segments {
		var name strings.Builder
		for _, word := range strings.Split(segment, "_") {
			if word == "" {
				continue
			}
			name.WriteString(strings.ToUpper(word[:1]))
			name.WriteString(word[1:])
		}
		segments[index] = name.String()
	}
	return strings.Join(segments, ".")
}

func normalizeGoPath(path string) string {
	segments := strings.Split(path, ".")
	for index, segment := range segments {
		segment = lastUpperBoundary.ReplaceAllString(segment, "$1 $2")
		segment = lowerUpperBoundary.ReplaceAllString(segment, "$1 $2")
		segment = nonNameCharacters.ReplaceAllString(segment, " ")
		segments[index] = strings.Join(strings.Fields(strings.ToLower(segment)), "_")
	}
	return strings.Join(segments, ".")
}

func collectClientMethods(root *sdkclient.Client) map[string]reflect.Method {
	methods := make(map[string]reflect.Method)
	var visit func(reflect.Value, string)
	visit = func(value reflect.Value, prefix string) {
		if value.Kind() != reflect.Pointer || value.IsNil() {
			return
		}
		clientType := value.Type()
		for index := 0; index < clientType.NumMethod(); index++ {
			method := clientType.Method(index)
			methods[prefix+method.Name] = method
		}

		clientValue := value.Elem()
		if clientValue.Kind() != reflect.Struct {
			return
		}
		for index := 0; index < clientValue.NumField(); index++ {
			field := clientValue.Type().Field(index)
			if field.PkgPath != "" || field.Type.Kind() != reflect.Pointer ||
				field.Type.Elem().Kind() != reflect.Struct || field.Type.Elem().Name() != "Client" {
				continue
			}
			visit(clientValue.Field(index), prefix+field.Name+".")
		}
	}
	visit(reflect.ValueOf(root), "")
	return methods
}

func findClientMethod(methods map[string]reflect.Method, operation string) (reflect.Method, string, bool) {
	plainName := goOperationName(operation)
	if method, found := methods[plainName]; found {
		return method, plainName, true
	}
	normalizedName := normalizeGoPath(operation)
	for actualName, method := range methods {
		if normalizeGoPath(actualName) == normalizedName {
			return method, actualName, true
		}
	}
	return reflect.Method{}, plainName, false
}

func requestOptionType(method reflect.Method) string {
	if !method.Type.IsVariadic() || method.Type.NumIn() == 0 {
		return ""
	}
	return method.Type.In(method.Type.NumIn() - 1).Elem().Name()
}

func isCorePageType(typeOf reflect.Type) bool {
	if typeOf.Kind() == reflect.Pointer {
		typeOf = typeOf.Elem()
	}
	pageType := reflect.TypeOf(core.Page[string, int, struct{}]{})
	return typeOf.PkgPath() == pageType.PkgPath() && strings.HasPrefix(typeOf.Name(), "Page[")
}

func sortedKeys(values map[string]bool) []string {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func TestClientMethodsMatchSection42(t *testing.T) {
	methods := collectClientMethods(newContractClient(nil, 1, option.WithAPIKey("contract-api-key")))
	for _, operation := range section42Operations {
		operation := operation
		t.Run(operation.name, func(t *testing.T) {
			method, actualName, found := findClientMethod(methods, operation.name)
			reason := missingOperationReason(operation.name)
			if !found {
				err := fmt.Errorf("client method %s is absent", goOperationName(operation.name))
				if reason != "" {
					expectContractFailure(t, reason, err)
					return
				}
				t.Fatal(err)
			}
			if reason != "" {
				expectContractFailure(t, reason, nil)
				return
			}
			if actualName != goOperationName(operation.name) {
				t.Logf("generated-name conversion exception: %s maps to %s", operation.name, actualName)
			}
			_ = method
		})
	}
}

func TestClientExposesNoMethodOutsideSection42(t *testing.T) {
	methods := collectClientMethods(newContractClient(nil, 1, option.WithAPIKey("contract-api-key")))
	expected := make(map[string]bool, len(section42Operations))
	for _, operation := range section42Operations {
		expected[normalizeGoPath(operation.name)] = true
	}

	var extras []string
	for methodName := range methods {
		if !expected[normalizeGoPath(methodName)] {
			extras = append(extras, methodName)
		}
	}
	sort.Strings(extras)
	if len(extras) > 0 {
		t.Errorf("client exposes methods outside section 4.2: %v", extras)
	}

	for _, operation := range sortedKeys(excludedFromSDK) {
		if _, _, found := findClientMethod(methods, operation); found {
			t.Errorf("excluded operation %s is present in the client", operation)
		}
	}
}

func TestPaginatedOperationsReturnPages(t *testing.T) {
	methods := collectClientMethods(newContractClient(nil, 1, option.WithAPIKey("contract-api-key")))
	for _, operation := range section42Operations {
		if !operation.paginated {
			continue
		}
		operation := operation
		t.Run(operation.name, func(t *testing.T) {
			method, _, found := findClientMethod(methods, operation.name)
			reason := missingOperationReason(operation.name)
			if !found {
				err := fmt.Errorf("client method %s is absent", goOperationName(operation.name))
				if reason != "" {
					expectContractFailure(t, reason, err)
					return
				}
				t.Fatal(err)
			}
			if method.Type.NumOut() == 0 || !isCorePageType(method.Type.Out(0)) {
				err := fmt.Errorf("%s does not return *core.Page", operation.name)
				if reason != "" {
					expectContractFailure(t, reason, err)
					return
				}
				t.Fatal(err)
			}
			if reason != "" {
				expectContractFailure(t, reason, nil)
			}
		})
	}
}

func TestOnlyStateChangingMethodsExposeIdempotencyKey(t *testing.T) {
	methods := collectClientMethods(newContractClient(nil, 1, option.WithAPIKey("contract-api-key")))
	requestOptionTypeName := reflect.TypeOf((*option.RequestOption)(nil)).Elem().Name()
	idempotentOptionTypeName := reflect.TypeOf((*option.IdempotentRequestOption)(nil)).Elem().Name()

	for _, operation := range section42Operations {
		operation := operation
		t.Run(operation.name, func(t *testing.T) {
			method, _, found := findClientMethod(methods, operation.name)
			var checkErr error
			if !found {
				checkErr = fmt.Errorf("client method %s is absent", goOperationName(operation.name))
			} else {
				expectedOption := requestOptionTypeName
				if operation.method != "GET" && !operation.postRead {
					expectedOption = idempotentOptionTypeName
				}
				actualOption := requestOptionType(method)
				if actualOption != expectedOption {
					checkErr = fmt.Errorf("%s accepts %q; want %q", operation.name, actualOption, expectedOption)
				}
			}

			reason := contractGapReason(operation.name, operation.postRead)
			if reason != "" {
				expectContractFailure(t, reason, checkErr)
				return
			}
			if checkErr != nil {
				t.Fatal(checkErr)
			}
		})
	}
}

type recordedRequest struct {
	method  string
	path    string
	cursor  string
	headers http.Header
}

type requestRecorder struct {
	mu       sync.Mutex
	requests []recordedRequest
}

func (recorder *requestRecorder) record(request *http.Request) int {
	recorder.mu.Lock()
	defer recorder.mu.Unlock()
	recorder.requests = append(recorder.requests, recordedRequest{
		method:  request.Method,
		path:    request.URL.Path,
		cursor:  request.URL.Query().Get("cursor"),
		headers: request.Header.Clone(),
	})
	return len(recorder.requests)
}

func (recorder *requestRecorder) snapshot() []recordedRequest {
	recorder.mu.Lock()
	defer recorder.mu.Unlock()
	return append([]recordedRequest(nil), recorder.requests...)
}

type contractResponseFunc func(http.ResponseWriter, *http.Request, int)

func newRecorderServer(t *testing.T, response contractResponseFunc) (*httptest.Server, *requestRecorder) {
	t.Helper()
	recorder := &requestRecorder{}
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		count := recorder.record(request)
		response(writer, request, count)
	}))
	t.Cleanup(server.Close)
	return server, recorder
}

func writeJSON(writer http.ResponseWriter, status int, value any) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	if value != nil {
		_ = json.NewEncoder(writer).Encode(value)
	}
}

func createUserRequest() *zep.CreateUserRequest {
	return &zep.CreateUserRequest{}
}

func TestPagerTraversesTwoPagesAndStopsWithoutNextCursor(t *testing.T) {
	for _, operation := range []string{"batch.list", "user.list"} {
		operation := operation
		t.Run(operation, func(t *testing.T) {
			path := "/batches"
			method := http.MethodGet
			if operation == "user.list" {
				path = "/users/list"
				method = http.MethodPost
			}
			server, recorder := newRecorderServer(t, func(writer http.ResponseWriter, request *http.Request, count int) {
				if request.Method != method || request.URL.Path != path {
					t.Errorf("received %s %s; want %s %s", request.Method, request.URL.Path, method, path)
				}
				nextCursor := ""
				if count == 1 {
					nextCursor = "c1"
				}
				items := []map[string]string{
					{"uuid": fmt.Sprintf("%s-%d", operation, 2*count-1)},
					{"uuid": fmt.Sprintf("%s-%d", operation, 2*count)},
				}
				body := map[string]any{"items": items}
				if nextCursor != "" {
					body["next_cursor"] = nextCursor
				}
				writeJSON(writer, http.StatusOK, body)
			})
			client := newContractClient(server, 1, option.WithAPIKey("contract-api-key"))
			ctx := context.Background()

			var items []string
			if operation == "batch.list" {
				page, err := client.Batch.List(ctx, &zep.BatchListRequest{Limit: zep.Int(2)})
				if err != nil {
					t.Fatalf("list batches: %v", err)
				}
				iterator := page.Iterator()
				for iterator.Next(ctx) {
					item := iterator.Current()
					if item.UUID == nil {
						t.Fatal("batch page item has no UUID")
					}
					items = append(items, *item.UUID)
				}
				if err := iterator.Err(); err != nil {
					t.Fatalf("iterate batches: %v", err)
				}
			} else {
				page, err := client.User.List(ctx, &zep.UserListRequest{Limit: zep.Int(2)})
				if err != nil {
					t.Fatalf("list users: %v", err)
				}
				iterator := page.Iterator()
				for iterator.Next(ctx) {
					item := iterator.Current()
					if item.UUID == nil {
						t.Fatal("user page item has no UUID")
					}
					items = append(items, *item.UUID)
				}
				if err := iterator.Err(); err != nil {
					t.Fatalf("iterate users: %v", err)
				}
			}

			requests := recorder.snapshot()
			if len(requests) != 2 {
				t.Fatalf("sent %d requests; want exactly two", len(requests))
			}
			if requests[1].cursor != "c1" {
				t.Fatalf("second request cursor is %q; want c1", requests[1].cursor)
			}
			wantItems := []string{
				fmt.Sprintf("%s-1", operation),
				fmt.Sprintf("%s-2", operation),
				fmt.Sprintf("%s-3", operation),
				fmt.Sprintf("%s-4", operation),
			}
			if !reflect.DeepEqual(items, wantItems) {
				t.Fatalf("items are %v; want %v", items, wantItems)
			}
			if len(items) != len(map[string]bool{
				items[0]: true,
				items[1]: true,
				items[2]: true,
				items[3]: true,
			}) {
				t.Fatalf("items are not distinct: %v", items)
			}
		})
	}
}

func TestCallWithoutKeySendsUUID4IdempotencyKey(t *testing.T) {
	server, recorder := newRecorderServer(t, func(writer http.ResponseWriter, _ *http.Request, _ int) {
		writeJSON(writer, http.StatusOK, map[string]any{})
	})
	client := newContractClient(server, 1, option.WithAPIKey("contract-api-key"))
	if _, err := client.User.Create(context.Background(), createUserRequest()); err != nil {
		t.Fatalf("create user: %v", err)
	}

	requests := recorder.snapshot()
	if len(requests) != 1 {
		t.Fatalf("sent %d requests; want one", len(requests))
	}
	key := requests[0].headers.Get("Idempotency-Key")
	if key == "" {
		expectContractFailure(t, d1Reason, fmt.Errorf("Idempotency-Key is absent"))
		return
	}
	parsed, err := uuid.Parse(key)
	if err != nil {
		t.Fatalf("Idempotency-Key %q is not a UUID: %v", key, err)
	}
	if parsed.Version() != uuid.Version(4) || parsed.Variant() != uuid.RFC4122 {
		t.Fatalf("Idempotency-Key %q is not an RFC 4122 UUIDv4", key)
	}
	expectContractFailure(t, d1Reason, nil)
}

func TestCallerIdempotencyKeyIsSentUnchanged(t *testing.T) {
	server, recorder := newRecorderServer(t, func(writer http.ResponseWriter, _ *http.Request, _ int) {
		writeJSON(writer, http.StatusOK, map[string]any{})
	})
	client := newContractClient(server, 1, option.WithAPIKey("contract-api-key"))
	if _, err := client.User.Create(
		context.Background(),
		createUserRequest(),
		option.WithIdempotencyKey(zep.String(callerKey)),
	); err != nil {
		t.Fatalf("create user: %v", err)
	}

	requests := recorder.snapshot()
	if len(requests) != 1 {
		t.Fatalf("sent %d requests; want one", len(requests))
	}
	actual := requests[0].headers.Get("Idempotency-Key")
	if actual == "*"+callerKey {
		expectContractFailure(t, d2Reason, fmt.Errorf("caller key was sent as %q", actual))
		return
	}
	if actual != callerKey {
		t.Fatalf("Idempotency-Key is %q; want %q", actual, callerKey)
	}
	expectContractFailure(t, d2Reason, nil)
}

func TestRetryReusesGeneratedIdempotencyKey(t *testing.T) {
	server, recorder := newRecorderServer(t, func(writer http.ResponseWriter, _ *http.Request, count int) {
		if count == 1 {
			writer.Header().Set("Retry-After", "1")
			writeJSON(writer, http.StatusServiceUnavailable, map[string]any{})
			return
		}
		writeJSON(writer, http.StatusOK, map[string]any{})
	})
	client := newContractClient(server, 2, option.WithAPIKey("contract-api-key"))
	if _, err := client.User.Create(context.Background(), createUserRequest()); err != nil {
		t.Fatalf("create user after retry: %v", err)
	}

	requests := recorder.snapshot()
	if len(requests) != 2 {
		t.Fatalf("sent %d requests; want exactly two", len(requests))
	}
	first := requests[0].headers.Get("Idempotency-Key")
	second := requests[1].headers.Get("Idempotency-Key")
	if first == "" && second == "" {
		expectContractFailure(t, d1Reason, fmt.Errorf("generated Idempotency-Key is absent on both attempts"))
		return
	}
	if first == "" || second == "" {
		t.Fatalf("Idempotency-Key differs across retries: first %q, second %q", first, second)
	}
	if first != second {
		t.Fatalf("Idempotency-Key changed across retries: first %q, second %q", first, second)
	}
	expectContractFailure(t, d1Reason, nil)
}

func TestRetryReusesCallerIdempotencyKey(t *testing.T) {
	server, recorder := newRecorderServer(t, func(writer http.ResponseWriter, _ *http.Request, count int) {
		if count == 1 {
			writer.Header().Set("Retry-After", "1")
			writeJSON(writer, http.StatusServiceUnavailable, map[string]any{})
			return
		}
		writeJSON(writer, http.StatusOK, map[string]any{})
	})
	client := newContractClient(server, 2, option.WithAPIKey("contract-api-key"))
	if _, err := client.User.Create(
		context.Background(),
		createUserRequest(),
		option.WithIdempotencyKey(zep.String(callerKey)),
	); err != nil {
		t.Fatalf("create user after retry: %v", err)
	}

	requests := recorder.snapshot()
	if len(requests) != 2 {
		t.Fatalf("sent %d requests; want exactly two", len(requests))
	}
	first := requests[0].headers.Get("Idempotency-Key")
	second := requests[1].headers.Get("Idempotency-Key")
	if first == "" {
		t.Fatal("Idempotency-Key is empty on the first request")
	}
	if first != second {
		t.Fatalf("Idempotency-Key changed across retries: first %q, second %q", first, second)
	}
}

func TestGETAndPOSTReadSendNoIdempotencyHeader(t *testing.T) {
	server, recorder := newRecorderServer(t, func(writer http.ResponseWriter, request *http.Request, _ int) {
		if request.Method == http.MethodPost {
			writeJSON(writer, http.StatusOK, map[string]any{"items": []any{}})
			return
		}
		writeJSON(writer, http.StatusOK, map[string]any{})
	})
	client := newContractClient(server, 1, option.WithAPIKey("contract-api-key"))
	ctx := context.Background()
	if _, err := client.Project.Get(ctx); err != nil {
		t.Fatalf("get project: %v", err)
	}
	if _, err := client.User.List(
		ctx,
		&zep.UserListRequest{Limit: zep.Int(1)},
	); err != nil {
		t.Fatalf("list users: %v", err)
	}

	requests := recorder.snapshot()
	if len(requests) != 2 {
		t.Fatalf("sent %d requests; want two", len(requests))
	}
	if requests[0].method != http.MethodGet || requests[0].path != "/project" {
		t.Fatalf("first request is %s %s; want GET /project", requests[0].method, requests[0].path)
	}
	if requests[0].headers.Get("Idempotency-Key") != "" {
		t.Fatalf("GET sent Idempotency-Key %q", requests[0].headers.Get("Idempotency-Key"))
	}
	if requests[1].method != http.MethodPost || requests[1].path != "/users/list" {
		t.Fatalf("second request is %s %s; want POST /users/list", requests[1].method, requests[1].path)
	}

	if key := requests[1].headers.Get("Idempotency-Key"); key != "" {
		t.Fatalf("POST read sent Idempotency-Key %q", key)
	}
}

func TestClientSendsProjectAPIKey(t *testing.T) {
	server, recorder := newRecorderServer(t, func(writer http.ResponseWriter, _ *http.Request, _ int) {
		writeJSON(writer, http.StatusOK, map[string]any{})
	})
	client := newContractClient(server, 1, option.WithAPIKey("contract-api-key"))
	if _, err := client.Project.Get(context.Background()); err != nil {
		t.Fatalf("get project: %v", err)
	}

	requests := recorder.snapshot()
	if len(requests) != 1 {
		t.Fatalf("sent %d requests; want one", len(requests))
	}
	if actual := requests[0].headers.Get("Authorization"); actual != "Api-Key contract-api-key" {
		t.Fatalf("Authorization is %q; want Api-Key contract-api-key", actual)
	}
	if actual := requests[0].headers.Get("X-Zep-Project"); actual != "" {
		t.Fatalf("X-Zep-Project is %q; want no project header", actual)
	}
}

func TestClientSendsAdminBearerWithProjectHeader(t *testing.T) {
	t.Setenv("ZEP_API_KEY", "")
	headers := make(http.Header)
	headers.Set("Authorization", "Bearer contract-token")
	headers.Set("X-Zep-Project", projectUUID)
	server, recorder := newRecorderServer(t, func(writer http.ResponseWriter, _ *http.Request, _ int) {
		writeJSON(writer, http.StatusOK, map[string]any{})
	})
	client := newContractClient(server, 1, option.WithHTTPHeader(headers))
	if _, err := client.Project.Get(context.Background()); err != nil {
		t.Fatalf("get project: %v", err)
	}

	requests := recorder.snapshot()
	if len(requests) != 1 {
		t.Fatalf("sent %d requests; want one", len(requests))
	}
	if actual := requests[0].headers.Get("Authorization"); actual != "Bearer contract-token" {
		t.Fatalf("Authorization is %q; want Bearer contract-token", actual)
	}
	if actual := requests[0].headers.Get("X-Zep-Project"); actual != projectUUID {
		t.Fatalf("X-Zep-Project is %q; want %s", actual, projectUUID)
	}
	if actual := requests[0].headers.Get("Api-Key"); actual != "" {
		t.Fatalf("Api-Key is %q; want no API-key header", actual)
	}
}

func stringConstantsForType(path string, typeName string) ([]string, error) {
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, path, nil, 0)
	if err != nil {
		return nil, err
	}
	var values []string
	for _, declaration := range file.Decls {
		general, ok := declaration.(*ast.GenDecl)
		if !ok || general.Tok != token.CONST {
			continue
		}
		currentType := ""
		for _, specification := range general.Specs {
			valueSpecification, ok := specification.(*ast.ValueSpec)
			if !ok {
				continue
			}
			if valueSpecification.Type != nil {
				if identifier, ok := valueSpecification.Type.(*ast.Ident); ok {
					currentType = identifier.Name
				} else {
					currentType = ""
				}
			}
			if currentType != typeName {
				continue
			}
			for _, expression := range valueSpecification.Values {
				literal, ok := expression.(*ast.BasicLit)
				if !ok || literal.Kind != token.STRING {
					return nil, fmt.Errorf("%s has a non-string constant for %s", path, typeName)
				}
				value, err := strconv.Unquote(literal.Value)
				if err != nil {
					return nil, fmt.Errorf("parse %s constant: %w", typeName, err)
				}
				values = append(values, value)
			}
		}
	}
	sort.Strings(values)
	return values, nil
}

func TestGraphContextRecencyBiasIsOffMildStrongStringEnum(t *testing.T) {
	requestType := reflect.TypeOf(zep.GraphContextRequest{})
	field, found := requestType.FieldByName("RecencyBias")
	if !found {
		t.Fatal("GraphContextRequest has no RecencyBias field")
	}
	if field.Type.Kind() != reflect.Pointer || field.Type.Elem().Kind() != reflect.String {
		t.Fatalf("GraphContextRequest.RecencyBias has type %s; want a pointer to a string enum", field.Type)
	}
	if field.Type.Elem().Name() != recencyBiasType {
		t.Fatalf("GraphContextRequest.RecencyBias uses %s; want %s", field.Type.Elem().Name(), recencyBiasType)
	}

	values, err := stringConstantsForType("graph.go", recencyBiasType)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"mild", "off", "strong"}
	if !reflect.DeepEqual(values, want) {
		t.Fatalf("%s values are %v; want exactly %v", recencyBiasType, values, want)
	}
}

func TestThreadContextRecencyBiasIsNotAnObject(t *testing.T) {
	graphType := reflect.TypeOf(zep.GraphContextRequest{})
	graphField, found := graphType.FieldByName("RecencyBias")
	if !found {
		t.Fatal("GraphContextRequest has no RecencyBias field")
	}

	threadType := reflect.TypeOf(zep.ThreadGetContextRequest{})
	threadField, found := threadType.FieldByName("RecencyBias")
	if found && threadField.Type != graphField.Type {
		t.Fatalf("ThreadGetContextRequest.RecencyBias has type %s; want the graph enum type %s or no field", threadField.Type, graphField.Type)
	}
}

func nameWords(name string) []string {
	name = lastUpperBoundary.ReplaceAllString(name, "$1 $2")
	name = lowerUpperBoundary.ReplaceAllString(name, "$1 $2")
	name = nonNameCharacters.ReplaceAllString(name, " ")
	return strings.Fields(strings.ToLower(name))
}

func v3OnlyConcept(name string) string {
	words := nameWords(name)
	for index, word := range words {
		if word == "scope" {
			return "scope"
		}
		if word == "lastn" || (word == "last" && index+1 < len(words) && words[index+1] == "n") {
			return "lastn"
		}
		if word == "uuid" && index+1 < len(words) && words[index+1] == "cursor" {
			return "uuid cursor"
		}
		if word == "group" && index+1 < len(words) && words[index+1] == "id" {
			return "group id"
		}
	}
	return ""
}

func publicGeneratedNameOffenders() ([]string, error) {
	fileSet := token.NewFileSet()
	var offenders []string
	err := filepath.WalkDir(".", func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			if entry.Name() == "internal" {
				return filepath.SkipDir
			}
			return nil
		}
		normalizedPath := strings.TrimPrefix(filepath.ToSlash(path), "./")
		if filepath.Ext(path) != ".go" || strings.HasSuffix(path, "_test.go") || normalizedPath == "ontology.go" {
			return nil
		}
		file, err := parser.ParseFile(fileSet, path, nil, 0)
		if err != nil {
			return err
		}
		add := func(name string, position token.Pos) {
			if concept := v3OnlyConcept(name); concept != "" {
				line := fileSet.Position(position).Line
				offenders = append(offenders, fmt.Sprintf("%s:%d: %s (%s)", path, line, name, concept))
			}
		}
		for _, declaration := range file.Decls {
			switch declaration := declaration.(type) {
			case *ast.GenDecl:
				if declaration.Tok != token.TYPE {
					continue
				}
				for _, specification := range declaration.Specs {
					if typeSpecification, ok := specification.(*ast.TypeSpec); ok && typeSpecification.Name.IsExported() {
						add(typeSpecification.Name.Name, typeSpecification.Name.Pos())
					}
				}
			case *ast.FuncDecl:
				if declaration.Name.IsExported() {
					add(declaration.Name.Name, declaration.Name.Pos())
				}
			}
		}
		return nil
	})
	sort.Strings(offenders)
	return offenders, err
}

func TestGeneratedNamesContainNoV3OnlyConcepts(t *testing.T) {
	offenders, err := publicGeneratedNameOffenders()
	if err != nil {
		t.Fatal(err)
	}
	const d4Reason = "The v4 Edge and ContextEdge schemas have a `scope` field (the edge kind), so the generated Go accessors GetScope and SetScope contain the v3-only word scope (spec 3 section 18.5)."
	known := map[string]bool{
		"ContextEdge.GetScope": true,
		"ContextEdge.SetScope": true,
		"Edge.GetScope":        true,
		"Edge.SetScope":        true,
	}
	found := make(map[string]bool, len(known))
	var unexpected []string
	for _, offender := range offenders {
		parts := strings.SplitN(offender, ":", 3)
		if len(parts) != 3 {
			unexpected = append(unexpected, offender)
			continue
		}
		methodName := strings.Fields(parts[2])
		if len(methodName) == 0 {
			unexpected = append(unexpected, offender)
			continue
		}
		file, err := parser.ParseFile(token.NewFileSet(), parts[0], nil, 0)
		if err != nil {
			unexpected = append(unexpected, fmt.Sprintf("%s: %v", offender, err))
			continue
		}
		matched := false
		for _, declaration := range file.Decls {
			method, ok := declaration.(*ast.FuncDecl)
			if !ok || method.Name.Name != methodName[0] || method.Recv == nil {
				continue
			}
			receiver := method.Recv.List[0].Type
			if pointer, ok := receiver.(*ast.StarExpr); ok {
				receiver = pointer.X
			}
			receiverType, ok := receiver.(*ast.Ident)
			if !ok {
				continue
			}
			name := receiverType.Name + "." + method.Name.Name
			matched = true
			if !known[name] {
				unexpected = append(unexpected, fmt.Sprintf("%s (%s)", offender, name))
				continue
			}
			found[name] = true
			t.Logf("open contract defect D4: %s: %s", name, d4Reason)
		}
		if !matched {
			unexpected = append(unexpected, offender)
		}
	}
	for name := range known {
		if !found[name] {
			t.Errorf("D4 is fixed for %s; remove the entry", name)
		}
	}
	if len(unexpected) > 0 {
		t.Errorf("public generated names contain unlisted v3-only concepts:\n%s", strings.Join(unexpected, "\n"))
	}
}
