// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/remoteworkers/v1test2/bots.proto

package remoteworkers // import "google.golang.org/genproto/googleapis/devtools/remoteworkers/v1test2"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"
import empty "github.com/golang/protobuf/ptypes/empty"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import status "google.golang.org/genproto/googleapis/rpc/status"
import field_mask "google.golang.org/genproto/protobuf/field_mask"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A coarse description of the status of the bot that the server uses to
// determine whether to assign the bot new leases.
type BotStatus int32

const (
	// Default value; do not use.
	BotStatus_BOT_STATUS_UNSPECIFIED BotStatus = 0
	// The bot is healthy, and will accept leases as normal.
	BotStatus_OK BotStatus = 1
	// The bot is unhealthy and will not accept new leases. For example, the bot
	// may have detected that available disk space is too low. This situation may
	// resolve itself, but will typically require human intervention.
	BotStatus_UNHEALTHY BotStatus = 2
	// The bot has been asked to reboot the host. The bot will not accept new
	// leases; once all leases are complete, this session will no longer be
	// updated but the bot will be expected to establish a new session after the
	// reboot completes.
	BotStatus_HOST_REBOOTING BotStatus = 3
	// The bot has been asked to shut down. As with HOST_REBOOTING, once all
	// leases are completed, the session will no longer be updated and the bot
	// will not be expected to establish a new session.
	//
	// Bots are typically only asked to shut down if its host computer will be
	// modified in some way, such as deleting a VM.
	BotStatus_BOT_TERMINATING BotStatus = 4
)

var BotStatus_name = map[int32]string{
	0: "BOT_STATUS_UNSPECIFIED",
	1: "OK",
	2: "UNHEALTHY",
	3: "HOST_REBOOTING",
	4: "BOT_TERMINATING",
}
var BotStatus_value = map[string]int32{
	"BOT_STATUS_UNSPECIFIED": 0,
	"OK":              1,
	"UNHEALTHY":       2,
	"HOST_REBOOTING":  3,
	"BOT_TERMINATING": 4,
}

func (x BotStatus) String() string {
	return proto.EnumName(BotStatus_name, int32(x))
}
func (BotStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bots_a82df967b734bf1b, []int{0}
}

// The state of the lease. All leases start in the PENDING state. A bot can
// change PENDING to ACTIVE or (in the case of an error) COMPLETED, or from
// ACTIVE to COMPLETED. The server can change PENDING or ACTIVE to CANCELLED if
// it wants the bot to release its resources - for example, if the bot needs to
// be quarantined (it's producing bad output) or a cell needs to be drained.
type LeaseState int32

const (
	// Default value; do not use.
	LeaseState_LEASE_STATE_UNSPECIFIED LeaseState = 0
	// Pending: the server expects the bot to accept this lease. This may only be
	// set by the server.
	LeaseState_PENDING LeaseState = 1
	// Active: the bot has accepted this lease. This may only be set by the bot.
	LeaseState_ACTIVE LeaseState = 2
	// Completed: the bot is no longer leased. This may only be set by the bot,
	// and the status field must be populated iff the state is COMPLETED.
	LeaseState_COMPLETED LeaseState = 4
	// Cancelled: The bot should immediately release all resources associated with
	// the lease. This may only be set by the server.
	LeaseState_CANCELLED LeaseState = 5
)

var LeaseState_name = map[int32]string{
	0: "LEASE_STATE_UNSPECIFIED",
	1: "PENDING",
	2: "ACTIVE",
	4: "COMPLETED",
	5: "CANCELLED",
}
var LeaseState_value = map[string]int32{
	"LEASE_STATE_UNSPECIFIED": 0,
	"PENDING":                 1,
	"ACTIVE":                  2,
	"COMPLETED":               4,
	"CANCELLED":               5,
}

func (x LeaseState) String() string {
	return proto.EnumName(LeaseState_name, int32(x))
}
func (LeaseState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bots_a82df967b734bf1b, []int{1}
}

// Possible administration actions.
type AdminTemp_Command int32

const (
	// Illegal value.
	AdminTemp_UNSPECIFIED AdminTemp_Command = 0
	// Download and run a new version of the bot. `arg` will be a resource
	// accessible via `ByteStream.Read` to obtain the new bot code.
	AdminTemp_BOT_UPDATE AdminTemp_Command = 1
	// Restart the bot without downloading a new version. `arg` will be a
	// message to log.
	AdminTemp_BOT_RESTART AdminTemp_Command = 2
	// Shut down the bot. `arg` will be a task resource name (similar to those
	// in tasks.proto) that the bot can use to tell the server that it is
	// terminating.
	AdminTemp_BOT_TERMINATE AdminTemp_Command = 3
	// Restart the host computer. `arg` will be a message to log.
	AdminTemp_HOST_RESTART AdminTemp_Command = 4
)

var AdminTemp_Command_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "BOT_UPDATE",
	2: "BOT_RESTART",
	3: "BOT_TERMINATE",
	4: "HOST_RESTART",
}
var AdminTemp_Command_value = map[string]int32{
	"UNSPECIFIED":   0,
	"BOT_UPDATE":    1,
	"BOT_RESTART":   2,
	"BOT_TERMINATE": 3,
	"HOST_RESTART":  4,
}

func (x AdminTemp_Command) String() string {
	return proto.EnumName(AdminTemp_Command_name, int32(x))
}
func (AdminTemp_Command) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bots_a82df967b734bf1b, []int{2, 0}
}

// Types of bot events.
type PostBotEventTempRequest_Type int32

const (
	// Illegal value.
	PostBotEventTempRequest_UNSPECIFIED PostBotEventTempRequest_Type = 0
	// Interesting but harmless event.
	PostBotEventTempRequest_INFO PostBotEventTempRequest_Type = 1
	// Error condition.
	PostBotEventTempRequest_ERROR PostBotEventTempRequest_Type = 2
)

var PostBotEventTempRequest_Type_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "INFO",
	2: "ERROR",
}
var PostBotEventTempRequest_Type_value = map[string]int32{
	"UNSPECIFIED": 0,
	"INFO":        1,
	"ERROR":       2,
}

func (x PostBotEventTempRequest_Type) String() string {
	return proto.EnumName(PostBotEventTempRequest_Type_name, int32(x))
}
func (PostBotEventTempRequest_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bots_a82df967b734bf1b, []int{5, 0}
}

// A bot session represents the state of a bot while in continuous contact with
// the server for a period of time. The session includes information about the
// worker - that is, the *worker* (the physical or virtual hardware) is
// considered to be a property of the bot (the software agent running on that
// hardware), which is the reverse of real life, but more natural from the point
// of the view of this API, which communicates solely with the bot and not
// directly with the underlying worker.
type BotSession struct {
	// The bot session name, as selected by the server. Output only during a call
	// to CreateBotSession.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A unique bot ID within the farm used to persistently identify this bot over
	// time (i.e., over multiple sessions). This ID must be unique within a
	// farm. Typically, the bot ID will be the same as the name of the primary
	// device in the worker (e.g., what you'd get from typing `uname -n` on *nix),
	// but this is not required since a single device may allow multiple bots to
	// run on it, each with access to different resources. What is important is
	// that this ID is meaningful to humans, who might need to hunt a physical
	// machine down to fix it.
	//
	// When CreateBotSession is successfully called with a bot_id, all prior
	// sessions with the same ID are invalidated. If a bot attempts to update an
	// invalid session, the server must reject that request, and may also
	// quarantine the other bot with the same bot IDs (ie, stop sending it new
	// leases and alert an admin).
	BotId string `protobuf:"bytes,2,opt,name=bot_id,json=botId,proto3" json:"bot_id,omitempty"`
	// The status of the bot. This must be populated in every call to
	// UpdateBotSession.
	Status BotStatus `protobuf:"varint,3,opt,name=status,proto3,enum=google.devtools.remoteworkers.v1test2.BotStatus" json:"status,omitempty"`
	// A description of the worker hosting this bot. The Worker message is used
	// here in the Status context (see Worker for more information).  If multiple
	// bots are running on the worker, this field should only describe the
	// resources accessible from this bot.
	//
	// During the call to CreateBotSession, the server may make arbitrary changes
	// to the worker's `server_properties` field (see that field for more
	// information). Otherwise, this field is input-only.
	Worker *Worker `protobuf:"bytes,4,opt,name=worker,proto3" json:"worker,omitempty"`
	// A list of all leases that are a part of this session. See the Lease message
	// for details.
	Leases []*Lease `protobuf:"bytes,5,rep,name=leases,proto3" json:"leases,omitempty"`
	// The time at which this bot session will expire, unless the bot calls
	// UpdateBotSession again. Output only.
	ExpireTime *timestamp.Timestamp `protobuf:"bytes,6,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
	// The version of the bot code currently running. The server may use this
	// information to issue an admin action to tell the bot to update itself.
	Version              string   `protobuf:"bytes,7,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BotSession) Reset()         { *m = BotSession{} }
func (m *BotSession) String() string { return proto.CompactTextString(m) }
func (*BotSession) ProtoMessage()    {}
func (*BotSession) Descriptor() ([]byte, []int) {
	return fileDescriptor_bots_a82df967b734bf1b, []int{0}
}
func (m *BotSession) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BotSession.Unmarshal(m, b)
}
func (m *BotSession) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BotSession.Marshal(b, m, deterministic)
}
func (dst *BotSession) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BotSession.Merge(dst, src)
}
func (m *BotSession) XXX_Size() int {
	return xxx_messageInfo_BotSession.Size(m)
}
func (m *BotSession) XXX_DiscardUnknown() {
	xxx_messageInfo_BotSession.DiscardUnknown(m)
}

var xxx_messageInfo_BotSession proto.InternalMessageInfo

func (m *BotSession) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BotSession) GetBotId() string {
	if m != nil {
		return m.BotId
	}
	return ""
}

func (m *BotSession) GetStatus() BotStatus {
	if m != nil {
		return m.Status
	}
	return BotStatus_BOT_STATUS_UNSPECIFIED
}

func (m *BotSession) GetWorker() *Worker {
	if m != nil {
		return m.Worker
	}
	return nil
}

func (m *BotSession) GetLeases() []*Lease {
	if m != nil {
		return m.Leases
	}
	return nil
}

func (m *BotSession) GetExpireTime() *timestamp.Timestamp {
	if m != nil {
		return m.ExpireTime
	}
	return nil
}

func (m *BotSession) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// A Lease is a lease that the scheduler has assigned to this bot. If the bot
// notices (by UpdateBotSession) that it has any leases in the PENDING state, it
// should call UpdateBotSession to put the leases into the ACTIVE state and
// start executing their assignments.
//
// All fields in this message are output-only, *except* the `state` and `status`
// fields. Note that repeated fields can only be updated as a unit, so on every
// update the bot must provide an update for *all* the leases the server expects
// it to report on.
//
// The scheduler *should* ensure that all leases scheduled to a bot can actually
// be accepted, but race conditions may occur. In such cases, the bot should
// attempt to accept the leases in the order they are listed by the server, to
// allow the server to control priorities.
//
// The server will remove COMPLETED leases from time to time, after which the
// bot shouldn't report on them any more (the server will ignore superfluous
// COMPLETED records).
type Lease struct {
	// The assignment, which is typically a resource that can be accessed through
	// some other services. The assignment must be meaningful to the bot based
	// solely on this name, either because the bot only understands one type of
	// assignment (e.g., tasks served through the Tasks interface) or through some
	// implementation-defined schema.
	//
	// For example, if the worker is executing a Task as defined by the Tasks
	// interface, this field might be projects/{projectid}/tasks/{taskid}.
	// However, it the worker is being assigned pull from a *queue* of tasks, the
	// resource would be the name of the queue, something like
	// projects/{projectid}/locations/{locationid}/queues/{queueid}. That queue
	// may then provide the bot with tasks through another interface, which the
	// bot then processes through the [Tasks]
	// [google.devtools.remoteworkers.v1test2.Tasks] interface.
	//
	// Note that the assignment may be a [full resource name]
	// [https://cloud.google.com/apis/design/resource_names#full_resource_name] if
	// it should be accessed through an endpoint that is not already known to the
	// bot.
	Assignment string `protobuf:"bytes,1,opt,name=assignment,proto3" json:"assignment,omitempty"`
	// The state of the lease. See LeaseState for more information.
	State LeaseState `protobuf:"varint,2,opt,name=state,proto3,enum=google.devtools.remoteworkers.v1test2.LeaseState" json:"state,omitempty"`
	// The final status of the lease (should be populated by the bot if the state
	// is completed). This is the status of the lease, not of any task represented
	// by the lease. For example, if the bot could not accept the lease because it
	// asked for some resource the bot didn't have, this status will be
	// FAILED_PRECONDITION. But if the assignment in the lease didn't execute
	// correctly, this field will be `OK` while the failure of the assignment must
	// be tracked elsewhere (e.g., through the Tasks interface).
	Status *status.Status `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	// The requirements that are being claimed by this lease. This field may be
	// omitted by the server if the lease is not pending.
	Requirements *Worker `protobuf:"bytes,4,opt,name=requirements,proto3" json:"requirements,omitempty"`
	// The time at which this lease expires. The server *may* extend this over
	// time, but due to race conditions, the bot is not *required* to respect any
	// expiry date except the first one.
	ExpireTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
	// While the `assignment` field is a resource name that allows the bot to
	// get the actual assignment, the server can also optionally include the
	// assignment itself inline in order to save a round trip to the server.
	//
	// This doesn't necessarily need to be the resource pointed to by
	// `assignment`. For example, if the assignment is a task, this field could
	// be task.description, not the entire task, since that's all the bot needs
	// to know to get started. As with `assignment` itself, all that's necessary
	// is that the bot knows how to handle the type of message received here.
	//
	// This field may be omitted by the server if the lease is not in the
	// `PENDING` state.
	InlineAssignment     *any.Any `protobuf:"bytes,6,opt,name=inline_assignment,json=inlineAssignment,proto3" json:"inline_assignment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Lease) Reset()         { *m = Lease{} }
func (m *Lease) String() string { return proto.CompactTextString(m) }
func (*Lease) ProtoMessage()    {}
func (*Lease) Descriptor() ([]byte, []int) {
	return fileDescriptor_bots_a82df967b734bf1b, []int{1}
}
func (m *Lease) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Lease.Unmarshal(m, b)
}
func (m *Lease) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Lease.Marshal(b, m, deterministic)
}
func (dst *Lease) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lease.Merge(dst, src)
}
func (m *Lease) XXX_Size() int {
	return xxx_messageInfo_Lease.Size(m)
}
func (m *Lease) XXX_DiscardUnknown() {
	xxx_messageInfo_Lease.DiscardUnknown(m)
}

var xxx_messageInfo_Lease proto.InternalMessageInfo

func (m *Lease) GetAssignment() string {
	if m != nil {
		return m.Assignment
	}
	return ""
}

func (m *Lease) GetState() LeaseState {
	if m != nil {
		return m.State
	}
	return LeaseState_LEASE_STATE_UNSPECIFIED
}

func (m *Lease) GetStatus() *status.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Lease) GetRequirements() *Worker {
	if m != nil {
		return m.Requirements
	}
	return nil
}

func (m *Lease) GetExpireTime() *timestamp.Timestamp {
	if m != nil {
		return m.ExpireTime
	}
	return nil
}

func (m *Lease) GetInlineAssignment() *any.Any {
	if m != nil {
		return m.InlineAssignment
	}
	return nil
}

// AdminTemp is a prelimiary set of administration tasks. It's called "Temp"
// because we do not yet know the best way to represent admin tasks; it's
// possible that this will be entirely replaced in later versions of this API.
// If this message proves to be sufficient, it will be renamed in the alpha or
// beta release of this API.
//
// This message (suitably marshalled into a protobuf.Any) can be used as the
// inline_assignment field in a lease; the lease assignment field should simply
// be `"admin"` in these cases.
//
// This message is heavily based on Swarming administration tasks from the LUCI
// project (http://github.com/luci/luci-py/appengine/swarming).
type AdminTemp struct {
	// The admin action; see `Command` for legal values.
	Command AdminTemp_Command `protobuf:"varint,1,opt,name=command,proto3,enum=google.devtools.remoteworkers.v1test2.AdminTemp_Command" json:"command,omitempty"`
	// The argument to the admin action; see `Command` for semantics.
	Arg                  string   `protobuf:"bytes,2,opt,name=arg,proto3" json:"arg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdminTemp) Reset()         { *m = AdminTemp{} }
func (m *AdminTemp) String() string { return proto.CompactTextString(m) }
func (*AdminTemp) ProtoMessage()    {}
func (*AdminTemp) Descriptor() ([]byte, []int) {
	return fileDescriptor_bots_a82df967b734bf1b, []int{2}
}
func (m *AdminTemp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdminTemp.Unmarshal(m, b)
}
func (m *AdminTemp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdminTemp.Marshal(b, m, deterministic)
}
func (dst *AdminTemp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdminTemp.Merge(dst, src)
}
func (m *AdminTemp) XXX_Size() int {
	return xxx_messageInfo_AdminTemp.Size(m)
}
func (m *AdminTemp) XXX_DiscardUnknown() {
	xxx_messageInfo_AdminTemp.DiscardUnknown(m)
}

var xxx_messageInfo_AdminTemp proto.InternalMessageInfo

func (m *AdminTemp) GetCommand() AdminTemp_Command {
	if m != nil {
		return m.Command
	}
	return AdminTemp_UNSPECIFIED
}

func (m *AdminTemp) GetArg() string {
	if m != nil {
		return m.Arg
	}
	return ""
}

// Request message for CreateBotSession.
type CreateBotSessionRequest struct {
	// The farm resource.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The bot session to create. Server-assigned fields like name must be unset.
	BotSession           *BotSession `protobuf:"bytes,2,opt,name=bot_session,json=botSession,proto3" json:"bot_session,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CreateBotSessionRequest) Reset()         { *m = CreateBotSessionRequest{} }
func (m *CreateBotSessionRequest) String() string { return proto.CompactTextString(m) }
func (*CreateBotSessionRequest) ProtoMessage()    {}
func (*CreateBotSessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bots_a82df967b734bf1b, []int{3}
}
func (m *CreateBotSessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBotSessionRequest.Unmarshal(m, b)
}
func (m *CreateBotSessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBotSessionRequest.Marshal(b, m, deterministic)
}
func (dst *CreateBotSessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBotSessionRequest.Merge(dst, src)
}
func (m *CreateBotSessionRequest) XXX_Size() int {
	return xxx_messageInfo_CreateBotSessionRequest.Size(m)
}
func (m *CreateBotSessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBotSessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBotSessionRequest proto.InternalMessageInfo

func (m *CreateBotSessionRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateBotSessionRequest) GetBotSession() *BotSession {
	if m != nil {
		return m.BotSession
	}
	return nil
}

// Request message for UpdateBotSession.
type UpdateBotSessionRequest struct {
	// The bot session name. Must match bot_session.name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The bot session resource to update.
	BotSession *BotSession `protobuf:"bytes,2,opt,name=bot_session,json=botSession,proto3" json:"bot_session,omitempty"`
	// The fields on the bot that should be updated. See the BotSession resource
	// for which fields are updatable by which caller.
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateBotSessionRequest) Reset()         { *m = UpdateBotSessionRequest{} }
func (m *UpdateBotSessionRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateBotSessionRequest) ProtoMessage()    {}
func (*UpdateBotSessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bots_a82df967b734bf1b, []int{4}
}
func (m *UpdateBotSessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBotSessionRequest.Unmarshal(m, b)
}
func (m *UpdateBotSessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBotSessionRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateBotSessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBotSessionRequest.Merge(dst, src)
}
func (m *UpdateBotSessionRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateBotSessionRequest.Size(m)
}
func (m *UpdateBotSessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBotSessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBotSessionRequest proto.InternalMessageInfo

func (m *UpdateBotSessionRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateBotSessionRequest) GetBotSession() *BotSession {
	if m != nil {
		return m.BotSession
	}
	return nil
}

func (m *UpdateBotSessionRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

// Request message for PostBotEventTemp
type PostBotEventTempRequest struct {
	// The bot session name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The type of bot event.
	Type PostBotEventTempRequest_Type `protobuf:"varint,2,opt,name=type,proto3,enum=google.devtools.remoteworkers.v1test2.PostBotEventTempRequest_Type" json:"type,omitempty"`
	// A human-readable message.
	Msg                  string   `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostBotEventTempRequest) Reset()         { *m = PostBotEventTempRequest{} }
func (m *PostBotEventTempRequest) String() string { return proto.CompactTextString(m) }
func (*PostBotEventTempRequest) ProtoMessage()    {}
func (*PostBotEventTempRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bots_a82df967b734bf1b, []int{5}
}
func (m *PostBotEventTempRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostBotEventTempRequest.Unmarshal(m, b)
}
func (m *PostBotEventTempRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostBotEventTempRequest.Marshal(b, m, deterministic)
}
func (dst *PostBotEventTempRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostBotEventTempRequest.Merge(dst, src)
}
func (m *PostBotEventTempRequest) XXX_Size() int {
	return xxx_messageInfo_PostBotEventTempRequest.Size(m)
}
func (m *PostBotEventTempRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostBotEventTempRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostBotEventTempRequest proto.InternalMessageInfo

func (m *PostBotEventTempRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PostBotEventTempRequest) GetType() PostBotEventTempRequest_Type {
	if m != nil {
		return m.Type
	}
	return PostBotEventTempRequest_UNSPECIFIED
}

func (m *PostBotEventTempRequest) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*BotSession)(nil), "google.devtools.remoteworkers.v1test2.BotSession")
	proto.RegisterType((*Lease)(nil), "google.devtools.remoteworkers.v1test2.Lease")
	proto.RegisterType((*AdminTemp)(nil), "google.devtools.remoteworkers.v1test2.AdminTemp")
	proto.RegisterType((*CreateBotSessionRequest)(nil), "google.devtools.remoteworkers.v1test2.CreateBotSessionRequest")
	proto.RegisterType((*UpdateBotSessionRequest)(nil), "google.devtools.remoteworkers.v1test2.UpdateBotSessionRequest")
	proto.RegisterType((*PostBotEventTempRequest)(nil), "google.devtools.remoteworkers.v1test2.PostBotEventTempRequest")
	proto.RegisterEnum("google.devtools.remoteworkers.v1test2.BotStatus", BotStatus_name, BotStatus_value)
	proto.RegisterEnum("google.devtools.remoteworkers.v1test2.LeaseState", LeaseState_name, LeaseState_value)
	proto.RegisterEnum("google.devtools.remoteworkers.v1test2.AdminTemp_Command", AdminTemp_Command_name, AdminTemp_Command_value)
	proto.RegisterEnum("google.devtools.remoteworkers.v1test2.PostBotEventTempRequest_Type", PostBotEventTempRequest_Type_name, PostBotEventTempRequest_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BotsClient is the client API for Bots service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BotsClient interface {
	// CreateBotSession is called when the bot first joins the farm, and
	// establishes a session ID to ensure that multiple machines do not register
	// using the same name accidentally.
	CreateBotSession(ctx context.Context, in *CreateBotSessionRequest, opts ...grpc.CallOption) (*BotSession, error)
	// UpdateBotSession must be called periodically by the bot (on a schedule
	// determined by the server) to let the server know about its status, and to
	// pick up new lease requests from the server.
	UpdateBotSession(ctx context.Context, in *UpdateBotSessionRequest, opts ...grpc.CallOption) (*BotSession, error)
	// PostBotEventTemp may be called by the bot to indicate that some exceptional
	// event has occurred. This method is subject to change or removal in future
	// revisions of this API; we may simply want to replace it with StackDriver or
	// some other common interface.
	PostBotEventTemp(ctx context.Context, in *PostBotEventTempRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type botsClient struct {
	cc *grpc.ClientConn
}

func NewBotsClient(cc *grpc.ClientConn) BotsClient {
	return &botsClient{cc}
}

func (c *botsClient) CreateBotSession(ctx context.Context, in *CreateBotSessionRequest, opts ...grpc.CallOption) (*BotSession, error) {
	out := new(BotSession)
	err := c.cc.Invoke(ctx, "/google.devtools.remoteworkers.v1test2.Bots/CreateBotSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botsClient) UpdateBotSession(ctx context.Context, in *UpdateBotSessionRequest, opts ...grpc.CallOption) (*BotSession, error) {
	out := new(BotSession)
	err := c.cc.Invoke(ctx, "/google.devtools.remoteworkers.v1test2.Bots/UpdateBotSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botsClient) PostBotEventTemp(ctx context.Context, in *PostBotEventTempRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.devtools.remoteworkers.v1test2.Bots/PostBotEventTemp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BotsServer is the server API for Bots service.
type BotsServer interface {
	// CreateBotSession is called when the bot first joins the farm, and
	// establishes a session ID to ensure that multiple machines do not register
	// using the same name accidentally.
	CreateBotSession(context.Context, *CreateBotSessionRequest) (*BotSession, error)
	// UpdateBotSession must be called periodically by the bot (on a schedule
	// determined by the server) to let the server know about its status, and to
	// pick up new lease requests from the server.
	UpdateBotSession(context.Context, *UpdateBotSessionRequest) (*BotSession, error)
	// PostBotEventTemp may be called by the bot to indicate that some exceptional
	// event has occurred. This method is subject to change or removal in future
	// revisions of this API; we may simply want to replace it with StackDriver or
	// some other common interface.
	PostBotEventTemp(context.Context, *PostBotEventTempRequest) (*empty.Empty, error)
}

func RegisterBotsServer(s *grpc.Server, srv BotsServer) {
	s.RegisterService(&_Bots_serviceDesc, srv)
}

func _Bots_CreateBotSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBotSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotsServer).CreateBotSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.remoteworkers.v1test2.Bots/CreateBotSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotsServer).CreateBotSession(ctx, req.(*CreateBotSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bots_UpdateBotSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBotSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotsServer).UpdateBotSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.remoteworkers.v1test2.Bots/UpdateBotSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotsServer).UpdateBotSession(ctx, req.(*UpdateBotSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bots_PostBotEventTemp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostBotEventTempRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotsServer).PostBotEventTemp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.remoteworkers.v1test2.Bots/PostBotEventTemp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotsServer).PostBotEventTemp(ctx, req.(*PostBotEventTempRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Bots_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.devtools.remoteworkers.v1test2.Bots",
	HandlerType: (*BotsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBotSession",
			Handler:    _Bots_CreateBotSession_Handler,
		},
		{
			MethodName: "UpdateBotSession",
			Handler:    _Bots_UpdateBotSession_Handler,
		},
		{
			MethodName: "PostBotEventTemp",
			Handler:    _Bots_PostBotEventTemp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/devtools/remoteworkers/v1test2/bots.proto",
}

func init() {
	proto.RegisterFile("google/devtools/remoteworkers/v1test2/bots.proto", fileDescriptor_bots_a82df967b734bf1b)
}

var fileDescriptor_bots_a82df967b734bf1b = []byte{
	// 1036 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xdf, 0x6e, 0xe3, 0xc4,
	0x17, 0xfe, 0xd9, 0x4d, 0xd2, 0x5f, 0x4e, 0x76, 0xbb, 0xee, 0x00, 0x6d, 0xc8, 0x22, 0x88, 0x22,
	0xad, 0x54, 0xa2, 0xc5, 0xde, 0x06, 0x21, 0xa1, 0x5d, 0x2d, 0x92, 0x93, 0xb8, 0x6d, 0x44, 0x9a,
	0x84, 0x89, 0xbb, 0x15, 0xdc, 0x64, 0x9d, 0x66, 0xd6, 0xb2, 0x1a, 0x7b, 0xbc, 0x9e, 0x69, 0xa1,
	0x42, 0x7b, 0x83, 0xc4, 0x13, 0xf0, 0x04, 0x70, 0x85, 0x78, 0x06, 0x84, 0xc4, 0x0d, 0x2f, 0xb0,
	0xaf, 0x00, 0xd7, 0xbc, 0x02, 0x9a, 0xf1, 0x24, 0x4d, 0xd2, 0x2d, 0xeb, 0x82, 0xb8, 0x9b, 0x3f,
	0xe7, 0x3b, 0xe7, 0x9b, 0x6f, 0xbe, 0x33, 0x36, 0x3c, 0xf0, 0x29, 0xf5, 0xa7, 0xc4, 0x9a, 0x90,
	0x73, 0x4e, 0xe9, 0x94, 0x59, 0x09, 0x09, 0x29, 0x27, 0x5f, 0xd2, 0xe4, 0x94, 0x24, 0xcc, 0x3a,
	0xdf, 0xe5, 0x84, 0xf1, 0x86, 0x35, 0xa6, 0x9c, 0x99, 0x71, 0x42, 0x39, 0x45, 0xf7, 0x52, 0x84,
	0x39, 0x43, 0x98, 0x4b, 0x08, 0x53, 0x21, 0x2a, 0xef, 0xa8, 0xc4, 0x5e, 0x1c, 0x58, 0x5e, 0x14,
	0x51, 0xee, 0xf1, 0x80, 0x46, 0x2a, 0x49, 0xa5, 0x91, 0xad, 0x6c, 0x3a, 0x57, 0x98, 0xb7, 0x15,
	0x46, 0xce, 0xc6, 0x67, 0xcf, 0x2c, 0x2f, 0xba, 0x50, 0x5b, 0x77, 0x57, 0xb7, 0x48, 0x18, 0xf3,
	0xd9, 0x66, 0x75, 0x75, 0xf3, 0x59, 0x40, 0xa6, 0x93, 0x51, 0xe8, 0xb1, 0x53, 0x15, 0xf1, 0xde,
	0x6a, 0x04, 0x0f, 0x42, 0xc2, 0xb8, 0x17, 0xc6, 0x2a, 0x60, 0x5b, 0x05, 0x24, 0xf1, 0x89, 0xc5,
	0xb8, 0xc7, 0xcf, 0xd4, 0x39, 0x6a, 0x7f, 0xe8, 0x00, 0x4d, 0xca, 0x87, 0x84, 0xb1, 0x80, 0x46,
	0x08, 0x41, 0x2e, 0xf2, 0x42, 0x52, 0xd6, 0xaa, 0xda, 0x4e, 0x11, 0xcb, 0x31, 0x7a, 0x0b, 0x0a,
	0x63, 0xca, 0x47, 0xc1, 0xa4, 0xac, 0xcb, 0xd5, 0xfc, 0x98, 0xf2, 0xce, 0x04, 0x1d, 0x40, 0x21,
	0xcd, 0x54, 0x5e, 0xab, 0x6a, 0x3b, 0x1b, 0x8d, 0x07, 0x66, 0x26, 0x5d, 0x4d, 0x51, 0x4d, 0xe2,
	0xb0, 0xc2, 0x23, 0x07, 0x0a, 0x69, 0x50, 0x39, 0x57, 0xd5, 0x76, 0x4a, 0x8d, 0x0f, 0x32, 0x66,
	0x3a, 0x96, 0x73, 0xac, 0xc0, 0xa8, 0x0d, 0x85, 0x29, 0xf1, 0x18, 0x61, 0xe5, 0x7c, 0x75, 0x6d,
	0xa7, 0xd4, 0xb8, 0x9f, 0x31, 0x4d, 0x57, 0x80, 0xb0, 0xc2, 0xa2, 0x47, 0x50, 0x22, 0x5f, 0xc5,
	0x41, 0x42, 0x46, 0x42, 0xc3, 0x72, 0x41, 0x32, 0xaa, 0xcc, 0x52, 0xcd, 0x04, 0x36, 0xdd, 0x99,
	0xc0, 0x18, 0xd2, 0x70, 0xb1, 0x80, 0xca, 0xb0, 0x7e, 0x4e, 0x12, 0xa1, 0x64, 0x79, 0x5d, 0x6a,
	0x35, 0x9b, 0xd6, 0xfe, 0xd4, 0x21, 0x2f, 0x0b, 0xa1, 0x77, 0x01, 0x3c, 0xc6, 0x02, 0x3f, 0x0a,
	0x49, 0xc4, 0x95, 0xd0, 0x0b, 0x2b, 0x68, 0x1f, 0xf2, 0x42, 0x17, 0x22, 0xd5, 0xde, 0x68, 0xec,
	0xde, 0xe4, 0x14, 0x42, 0x58, 0x82, 0x53, 0x3c, 0xaa, 0x2f, 0x5d, 0x50, 0xa9, 0x81, 0x66, 0x99,
	0x92, 0xf8, 0xc4, 0x5c, 0xb9, 0x82, 0xcf, 0xe0, 0x56, 0x42, 0x9e, 0x9f, 0x05, 0x09, 0x11, 0x1c,
	0xd8, 0x3f, 0xbb, 0x88, 0xa5, 0x14, 0xab, 0x42, 0xe6, 0x6f, 0x24, 0xa4, 0x0d, 0x9b, 0x41, 0x34,
	0x0d, 0x22, 0x32, 0x5a, 0xd0, 0x2a, 0xbd, 0x8b, 0x37, 0xaf, 0xa4, 0xb0, 0xa3, 0x0b, 0x6c, 0xa4,
	0xe1, 0xf6, 0x3c, 0xba, 0xf6, 0x52, 0x83, 0xa2, 0x3d, 0x09, 0x83, 0xc8, 0x25, 0x61, 0x8c, 0x30,
	0xac, 0x9f, 0xd0, 0x30, 0xf4, 0xa2, 0x89, 0x94, 0x7c, 0xa3, 0xf1, 0x71, 0xc6, 0xb3, 0xcd, 0x53,
	0x98, 0xad, 0x14, 0x8f, 0x67, 0x89, 0x90, 0x01, 0x6b, 0x5e, 0xe2, 0xab, 0xae, 0x10, 0xc3, 0xda,
	0x53, 0x58, 0x57, 0x51, 0xe8, 0x0e, 0x94, 0x8e, 0x7a, 0xc3, 0x81, 0xd3, 0xea, 0xec, 0x75, 0x9c,
	0xb6, 0xf1, 0x3f, 0xb4, 0x01, 0xd0, 0xec, 0xbb, 0xa3, 0xa3, 0x41, 0xdb, 0x76, 0x1d, 0x43, 0x13,
	0x01, 0x62, 0x8e, 0x9d, 0xa1, 0x6b, 0x63, 0xd7, 0xd0, 0xd1, 0x26, 0xdc, 0x16, 0x0b, 0xae, 0x83,
	0x0f, 0x3b, 0x3d, 0x11, 0xb3, 0x86, 0x0c, 0xb8, 0x75, 0xd0, 0x1f, 0x5e, 0x06, 0xe5, 0x6a, 0xdf,
	0x6a, 0xb0, 0xdd, 0x4a, 0x88, 0xc7, 0xc9, 0x65, 0xd7, 0x62, 0xf2, 0xfc, 0x8c, 0x30, 0x8e, 0xb6,
	0xa0, 0x10, 0x7b, 0xc9, 0xa5, 0xab, 0xd4, 0x0c, 0x61, 0x28, 0x89, 0x06, 0x66, 0x69, 0xb4, 0xe4,
	0x5b, 0xca, 0xec, 0xab, 0x85, 0x32, 0x30, 0x9e, 0x8f, 0x6b, 0xbf, 0x68, 0xb0, 0x7d, 0x14, 0x4f,
	0x5e, 0xc9, 0xe3, 0x55, 0x8f, 0xc8, 0x7f, 0xc0, 0x41, 0x38, 0xec, 0x4c, 0x52, 0x90, 0x4f, 0xa1,
	0x72, 0xf9, 0x55, 0x87, 0xed, 0x89, 0xd7, 0xf2, 0xd0, 0x63, 0xa7, 0x18, 0xd2, 0x70, 0x31, 0xae,
	0xfd, 0xa6, 0xc1, 0xf6, 0x80, 0x32, 0xde, 0xa4, 0xdc, 0x39, 0x27, 0x11, 0x17, 0x57, 0xfc, 0x77,
	0x07, 0x38, 0x86, 0x1c, 0xbf, 0x88, 0x67, 0x5d, 0xd9, 0xca, 0xc8, 0xfc, 0x9a, 0x0a, 0xa6, 0x7b,
	0x11, 0x13, 0x2c, 0x13, 0x0a, 0x17, 0x85, 0xcc, 0x97, 0xec, 0x8b, 0x58, 0x0c, 0x6b, 0xf7, 0x21,
	0x27, 0xf6, 0xaf, 0x5a, 0xe8, 0xff, 0x90, 0xeb, 0xf4, 0xf6, 0xfa, 0x86, 0x86, 0x8a, 0x90, 0x77,
	0x30, 0xee, 0x63, 0x43, 0xaf, 0xfb, 0x50, 0x9c, 0x3f, 0xa9, 0xa8, 0x02, 0x5b, 0xc2, 0x43, 0x43,
	0xd7, 0x76, 0x8f, 0x86, 0xa3, 0x65, 0x74, 0x01, 0xf4, 0xfe, 0xa7, 0x86, 0x86, 0x6e, 0x43, 0xf1,
	0xa8, 0x77, 0xe0, 0xd8, 0x5d, 0xf7, 0xe0, 0x73, 0x43, 0x47, 0x08, 0x36, 0x94, 0xc7, 0x9a, 0xfd,
	0xbe, 0xdb, 0xe9, 0xed, 0x1b, 0x6b, 0xe8, 0x0d, 0xb8, 0xb3, 0x68, 0x45, 0xb1, 0x98, 0xab, 0x3f,
	0x05, 0xb8, 0x7c, 0x64, 0xd0, 0x5d, 0xd8, 0xee, 0x3a, 0xf6, 0xd0, 0x91, 0xb5, 0x9c, 0x95, 0x52,
	0x25, 0x58, 0x1f, 0x38, 0xbd, 0xb6, 0xc0, 0x69, 0x08, 0xa0, 0x60, 0xb7, 0xdc, 0xce, 0x13, 0xc7,
	0xd0, 0x45, 0xed, 0x56, 0xff, 0x70, 0xd0, 0x75, 0x5c, 0xa7, 0x6d, 0xe4, 0xe4, 0xd4, 0xee, 0xb5,
	0x9c, 0x6e, 0xd7, 0x69, 0x1b, 0xf9, 0xc6, 0x8f, 0x39, 0xc8, 0x35, 0x29, 0x67, 0xe8, 0x67, 0x0d,
	0x8c, 0x55, 0x97, 0xa3, 0x4f, 0x32, 0x6a, 0x7e, 0x4d, 0x7b, 0x54, 0x6e, 0xee, 0xb6, 0xda, 0x47,
	0xdf, 0xbc, 0xfc, 0xfd, 0x3b, 0xdd, 0xaa, 0x55, 0xe7, 0x1f, 0xf4, 0xaf, 0xd3, 0x9e, 0x7a, 0x5c,
	0xaf, 0xbf, 0xb0, 0x2e, 0xad, 0xc8, 0x1e, 0x2e, 0xba, 0x5b, 0xd2, 0x5f, 0x6d, 0x8e, 0xcc, 0xf4,
	0xaf, 0xe9, 0xaa, 0x7f, 0x41, 0xbf, 0xb1, 0x40, 0x5f, 0x78, 0xf9, 0x71, 0xbd, 0xbe, 0xc8, 0xdd,
	0xaa, 0xbf, 0x58, 0xa6, 0xff, 0xbd, 0x06, 0xc6, 0xaa, 0x71, 0x33, 0xd3, 0xbf, 0xc6, 0xf1, 0x95,
	0xad, 0x2b, 0x7d, 0xe9, 0x88, 0x5f, 0x9c, 0xb9, 0xc4, 0xf5, 0xd7, 0x72, 0x8c, 0x29, 0x4b, 0xd3,
	0x3e, 0xd4, 0xea, 0xcd, 0x5f, 0x35, 0x78, 0xff, 0x84, 0x86, 0xd9, 0x48, 0x35, 0x37, 0xb1, 0x5c,
	0x4e, 0xbf, 0x53, 0x4c, 0x58, 0x6c, 0xa0, 0x7d, 0x81, 0x15, 0xd6, 0xa7, 0x53, 0x2f, 0xf2, 0x4d,
	0x9a, 0xf8, 0x96, 0x4f, 0x22, 0x49, 0xcf, 0x4a, 0xb7, 0xbc, 0x38, 0x60, 0xaf, 0xf9, 0xc3, 0x7b,
	0xb4, 0xb4, 0xfa, 0x83, 0xae, 0xe3, 0xe3, 0x9f, 0xf4, 0x7b, 0xfb, 0x69, 0xe6, 0x36, 0x39, 0x77,
	0x25, 0xab, 0xa5, 0xf2, 0xe6, 0x93, 0x5d, 0x57, 0x40, 0xc7, 0x05, 0x59, 0xeb, 0xc3, 0xbf, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x36, 0x3b, 0x58, 0x77, 0xc3, 0x0a, 0x00, 0x00,
}
