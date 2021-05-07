// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/scionproto/scion/go/pkg/gateway/control (interfaces: DataplaneSession,Discoverer,RoutingTable,RoutingTableSwapper,RoutingTableFactory,EngineFactory,PathMonitor,PathMonitorRegistration,PacketConnFactory,PrefixConsumer,PrefixFetcher,DataplaneSessionFactory,PktWriter,Worker,SessionPolicyParser,RoutingPolicyProvider,Runner,GatewayWatcherFactory,Publisher,PublisherFactory)

// Package mock_control is a generated GoMock package.
package mock_control

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	gopacket "github.com/google/gopacket"
	layers "github.com/google/gopacket/layers"
	addr "github.com/scionproto/scion/go/lib/addr"
	snet "github.com/scionproto/scion/go/lib/snet"
	control "github.com/scionproto/scion/go/pkg/gateway/control"
	pathhealth "github.com/scionproto/scion/go/pkg/gateway/pathhealth"
	policies "github.com/scionproto/scion/go/pkg/gateway/pathhealth/policies"
	routing "github.com/scionproto/scion/go/pkg/gateway/routing"
	net "net"
	reflect "reflect"
)

// MockDataplaneSession is a mock of DataplaneSession interface
type MockDataplaneSession struct {
	ctrl     *gomock.Controller
	recorder *MockDataplaneSessionMockRecorder
}

// MockDataplaneSessionMockRecorder is the mock recorder for MockDataplaneSession
type MockDataplaneSessionMockRecorder struct {
	mock *MockDataplaneSession
}

// NewMockDataplaneSession creates a new mock instance
func NewMockDataplaneSession(ctrl *gomock.Controller) *MockDataplaneSession {
	mock := &MockDataplaneSession{ctrl: ctrl}
	mock.recorder = &MockDataplaneSessionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDataplaneSession) EXPECT() *MockDataplaneSessionMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockDataplaneSession) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockDataplaneSessionMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDataplaneSession)(nil).Close))
}

// SetPaths mocks base method
func (m *MockDataplaneSession) SetPaths(arg0 []snet.Path) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPaths", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPaths indicates an expected call of SetPaths
func (mr *MockDataplaneSessionMockRecorder) SetPaths(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPaths", reflect.TypeOf((*MockDataplaneSession)(nil).SetPaths), arg0)
}

// Write mocks base method
func (m *MockDataplaneSession) Write(arg0 gopacket.Packet) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Write", arg0)
}

// Write indicates an expected call of Write
func (mr *MockDataplaneSessionMockRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockDataplaneSession)(nil).Write), arg0)
}

// MockDiscoverer is a mock of Discoverer interface
type MockDiscoverer struct {
	ctrl     *gomock.Controller
	recorder *MockDiscovererMockRecorder
}

// MockDiscovererMockRecorder is the mock recorder for MockDiscoverer
type MockDiscovererMockRecorder struct {
	mock *MockDiscoverer
}

// NewMockDiscoverer creates a new mock instance
func NewMockDiscoverer(ctrl *gomock.Controller) *MockDiscoverer {
	mock := &MockDiscoverer{ctrl: ctrl}
	mock.recorder = &MockDiscovererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDiscoverer) EXPECT() *MockDiscovererMockRecorder {
	return m.recorder
}

// Gateways mocks base method
func (m *MockDiscoverer) Gateways(arg0 context.Context) ([]control.Gateway, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Gateways", arg0)
	ret0, _ := ret[0].([]control.Gateway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Gateways indicates an expected call of Gateways
func (mr *MockDiscovererMockRecorder) Gateways(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Gateways", reflect.TypeOf((*MockDiscoverer)(nil).Gateways), arg0)
}

// MockRoutingTable is a mock of RoutingTable interface
type MockRoutingTable struct {
	ctrl     *gomock.Controller
	recorder *MockRoutingTableMockRecorder
}

// MockRoutingTableMockRecorder is the mock recorder for MockRoutingTable
type MockRoutingTableMockRecorder struct {
	mock *MockRoutingTable
}

// NewMockRoutingTable creates a new mock instance
func NewMockRoutingTable(ctrl *gomock.Controller) *MockRoutingTable {
	mock := &MockRoutingTable{ctrl: ctrl}
	mock.recorder = &MockRoutingTableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRoutingTable) EXPECT() *MockRoutingTableMockRecorder {
	return m.recorder
}

// Activate mocks base method
func (m *MockRoutingTable) Activate() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Activate")
}

// Activate indicates an expected call of Activate
func (mr *MockRoutingTableMockRecorder) Activate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Activate", reflect.TypeOf((*MockRoutingTable)(nil).Activate))
}

// ClearSession mocks base method
func (m *MockRoutingTable) ClearSession(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearSession", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearSession indicates an expected call of ClearSession
func (mr *MockRoutingTableMockRecorder) ClearSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearSession", reflect.TypeOf((*MockRoutingTable)(nil).ClearSession), arg0)
}

// Deactivate mocks base method
func (m *MockRoutingTable) Deactivate() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Deactivate")
}

// Deactivate indicates an expected call of Deactivate
func (mr *MockRoutingTableMockRecorder) Deactivate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deactivate", reflect.TypeOf((*MockRoutingTable)(nil).Deactivate))
}

// RouteIPv4 mocks base method
func (m *MockRoutingTable) RouteIPv4(arg0 layers.IPv4) control.PktWriter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RouteIPv4", arg0)
	ret0, _ := ret[0].(control.PktWriter)
	return ret0
}

// RouteIPv4 indicates an expected call of RouteIPv4
func (mr *MockRoutingTableMockRecorder) RouteIPv4(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteIPv4", reflect.TypeOf((*MockRoutingTable)(nil).RouteIPv4), arg0)
}

// RouteIPv6 mocks base method
func (m *MockRoutingTable) RouteIPv6(arg0 layers.IPv6) control.PktWriter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RouteIPv6", arg0)
	ret0, _ := ret[0].(control.PktWriter)
	return ret0
}

// RouteIPv6 indicates an expected call of RouteIPv6
func (mr *MockRoutingTableMockRecorder) RouteIPv6(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteIPv6", reflect.TypeOf((*MockRoutingTable)(nil).RouteIPv6), arg0)
}

// SetSession mocks base method
func (m *MockRoutingTable) SetSession(arg0 int, arg1 control.PktWriter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSession", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSession indicates an expected call of SetSession
func (mr *MockRoutingTableMockRecorder) SetSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSession", reflect.TypeOf((*MockRoutingTable)(nil).SetSession), arg0, arg1)
}

// MockRoutingTableSwapper is a mock of RoutingTableSwapper interface
type MockRoutingTableSwapper struct {
	ctrl     *gomock.Controller
	recorder *MockRoutingTableSwapperMockRecorder
}

// MockRoutingTableSwapperMockRecorder is the mock recorder for MockRoutingTableSwapper
type MockRoutingTableSwapperMockRecorder struct {
	mock *MockRoutingTableSwapper
}

// NewMockRoutingTableSwapper creates a new mock instance
func NewMockRoutingTableSwapper(ctrl *gomock.Controller) *MockRoutingTableSwapper {
	mock := &MockRoutingTableSwapper{ctrl: ctrl}
	mock.recorder = &MockRoutingTableSwapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRoutingTableSwapper) EXPECT() *MockRoutingTableSwapperMockRecorder {
	return m.recorder
}

// SetRoutingTable mocks base method
func (m *MockRoutingTableSwapper) SetRoutingTable(arg0 control.RoutingTable) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRoutingTable", arg0)
}

// SetRoutingTable indicates an expected call of SetRoutingTable
func (mr *MockRoutingTableSwapperMockRecorder) SetRoutingTable(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRoutingTable", reflect.TypeOf((*MockRoutingTableSwapper)(nil).SetRoutingTable), arg0)
}

// MockRoutingTableFactory is a mock of RoutingTableFactory interface
type MockRoutingTableFactory struct {
	ctrl     *gomock.Controller
	recorder *MockRoutingTableFactoryMockRecorder
}

// MockRoutingTableFactoryMockRecorder is the mock recorder for MockRoutingTableFactory
type MockRoutingTableFactoryMockRecorder struct {
	mock *MockRoutingTableFactory
}

// NewMockRoutingTableFactory creates a new mock instance
func NewMockRoutingTableFactory(ctrl *gomock.Controller) *MockRoutingTableFactory {
	mock := &MockRoutingTableFactory{ctrl: ctrl}
	mock.recorder = &MockRoutingTableFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRoutingTableFactory) EXPECT() *MockRoutingTableFactoryMockRecorder {
	return m.recorder
}

// New mocks base method
func (m *MockRoutingTableFactory) New(arg0 []*control.RoutingChain) (control.RoutingTable, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", arg0)
	ret0, _ := ret[0].(control.RoutingTable)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// New indicates an expected call of New
func (mr *MockRoutingTableFactoryMockRecorder) New(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockRoutingTableFactory)(nil).New), arg0)
}

// MockEngineFactory is a mock of EngineFactory interface
type MockEngineFactory struct {
	ctrl     *gomock.Controller
	recorder *MockEngineFactoryMockRecorder
}

// MockEngineFactoryMockRecorder is the mock recorder for MockEngineFactory
type MockEngineFactoryMockRecorder struct {
	mock *MockEngineFactory
}

// NewMockEngineFactory creates a new mock instance
func NewMockEngineFactory(ctrl *gomock.Controller) *MockEngineFactory {
	mock := &MockEngineFactory{ctrl: ctrl}
	mock.recorder = &MockEngineFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEngineFactory) EXPECT() *MockEngineFactoryMockRecorder {
	return m.recorder
}

// New mocks base method
func (m *MockEngineFactory) New(arg0 control.RoutingTable, arg1 []*control.SessionConfig, arg2 map[int][]byte) control.Worker {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", arg0, arg1, arg2)
	ret0, _ := ret[0].(control.Worker)
	return ret0
}

// New indicates an expected call of New
func (mr *MockEngineFactoryMockRecorder) New(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockEngineFactory)(nil).New), arg0, arg1, arg2)
}

// MockPathMonitor is a mock of PathMonitor interface
type MockPathMonitor struct {
	ctrl     *gomock.Controller
	recorder *MockPathMonitorMockRecorder
}

// MockPathMonitorMockRecorder is the mock recorder for MockPathMonitor
type MockPathMonitorMockRecorder struct {
	mock *MockPathMonitor
}

// NewMockPathMonitor creates a new mock instance
func NewMockPathMonitor(ctrl *gomock.Controller) *MockPathMonitor {
	mock := &MockPathMonitor{ctrl: ctrl}
	mock.recorder = &MockPathMonitorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPathMonitor) EXPECT() *MockPathMonitorMockRecorder {
	return m.recorder
}

// Register mocks base method
func (m *MockPathMonitor) Register(arg0 addr.IA, arg1 *policies.Policies, arg2 int) control.PathMonitorRegistration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1, arg2)
	ret0, _ := ret[0].(control.PathMonitorRegistration)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockPathMonitorMockRecorder) Register(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockPathMonitor)(nil).Register), arg0, arg1, arg2)
}

// MockPathMonitorRegistration is a mock of PathMonitorRegistration interface
type MockPathMonitorRegistration struct {
	ctrl     *gomock.Controller
	recorder *MockPathMonitorRegistrationMockRecorder
}

// MockPathMonitorRegistrationMockRecorder is the mock recorder for MockPathMonitorRegistration
type MockPathMonitorRegistrationMockRecorder struct {
	mock *MockPathMonitorRegistration
}

// NewMockPathMonitorRegistration creates a new mock instance
func NewMockPathMonitorRegistration(ctrl *gomock.Controller) *MockPathMonitorRegistration {
	mock := &MockPathMonitorRegistration{ctrl: ctrl}
	mock.recorder = &MockPathMonitorRegistrationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPathMonitorRegistration) EXPECT() *MockPathMonitorRegistrationMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockPathMonitorRegistration) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockPathMonitorRegistrationMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPathMonitorRegistration)(nil).Close))
}

// Get mocks base method
func (m *MockPathMonitorRegistration) Get() pathhealth.Selection {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(pathhealth.Selection)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockPathMonitorRegistrationMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPathMonitorRegistration)(nil).Get))
}

// MockPacketConnFactory is a mock of PacketConnFactory interface
type MockPacketConnFactory struct {
	ctrl     *gomock.Controller
	recorder *MockPacketConnFactoryMockRecorder
}

// MockPacketConnFactoryMockRecorder is the mock recorder for MockPacketConnFactory
type MockPacketConnFactoryMockRecorder struct {
	mock *MockPacketConnFactory
}

// NewMockPacketConnFactory creates a new mock instance
func NewMockPacketConnFactory(ctrl *gomock.Controller) *MockPacketConnFactory {
	mock := &MockPacketConnFactory{ctrl: ctrl}
	mock.recorder = &MockPacketConnFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPacketConnFactory) EXPECT() *MockPacketConnFactoryMockRecorder {
	return m.recorder
}

// New mocks base method
func (m *MockPacketConnFactory) New() (net.PacketConn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New")
	ret0, _ := ret[0].(net.PacketConn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// New indicates an expected call of New
func (mr *MockPacketConnFactoryMockRecorder) New() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockPacketConnFactory)(nil).New))
}

// MockPrefixConsumer is a mock of PrefixConsumer interface
type MockPrefixConsumer struct {
	ctrl     *gomock.Controller
	recorder *MockPrefixConsumerMockRecorder
}

// MockPrefixConsumerMockRecorder is the mock recorder for MockPrefixConsumer
type MockPrefixConsumerMockRecorder struct {
	mock *MockPrefixConsumer
}

// NewMockPrefixConsumer creates a new mock instance
func NewMockPrefixConsumer(ctrl *gomock.Controller) *MockPrefixConsumer {
	mock := &MockPrefixConsumer{ctrl: ctrl}
	mock.recorder = &MockPrefixConsumerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPrefixConsumer) EXPECT() *MockPrefixConsumerMockRecorder {
	return m.recorder
}

// Prefixes mocks base method
func (m *MockPrefixConsumer) Prefixes(arg0 addr.IA, arg1 control.Gateway, arg2 []*net.IPNet) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Prefixes", arg0, arg1, arg2)
}

// Prefixes indicates an expected call of Prefixes
func (mr *MockPrefixConsumerMockRecorder) Prefixes(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prefixes", reflect.TypeOf((*MockPrefixConsumer)(nil).Prefixes), arg0, arg1, arg2)
}

// MockPrefixFetcher is a mock of PrefixFetcher interface
type MockPrefixFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockPrefixFetcherMockRecorder
}

// MockPrefixFetcherMockRecorder is the mock recorder for MockPrefixFetcher
type MockPrefixFetcherMockRecorder struct {
	mock *MockPrefixFetcher
}

// NewMockPrefixFetcher creates a new mock instance
func NewMockPrefixFetcher(ctrl *gomock.Controller) *MockPrefixFetcher {
	mock := &MockPrefixFetcher{ctrl: ctrl}
	mock.recorder = &MockPrefixFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPrefixFetcher) EXPECT() *MockPrefixFetcherMockRecorder {
	return m.recorder
}

// Prefixes mocks base method
func (m *MockPrefixFetcher) Prefixes(arg0 context.Context, arg1 *net.UDPAddr) ([]*net.IPNet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prefixes", arg0, arg1)
	ret0, _ := ret[0].([]*net.IPNet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Prefixes indicates an expected call of Prefixes
func (mr *MockPrefixFetcherMockRecorder) Prefixes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prefixes", reflect.TypeOf((*MockPrefixFetcher)(nil).Prefixes), arg0, arg1)
}

// MockDataplaneSessionFactory is a mock of DataplaneSessionFactory interface
type MockDataplaneSessionFactory struct {
	ctrl     *gomock.Controller
	recorder *MockDataplaneSessionFactoryMockRecorder
}

// MockDataplaneSessionFactoryMockRecorder is the mock recorder for MockDataplaneSessionFactory
type MockDataplaneSessionFactoryMockRecorder struct {
	mock *MockDataplaneSessionFactory
}

// NewMockDataplaneSessionFactory creates a new mock instance
func NewMockDataplaneSessionFactory(ctrl *gomock.Controller) *MockDataplaneSessionFactory {
	mock := &MockDataplaneSessionFactory{ctrl: ctrl}
	mock.recorder = &MockDataplaneSessionFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDataplaneSessionFactory) EXPECT() *MockDataplaneSessionFactoryMockRecorder {
	return m.recorder
}

// New mocks base method
func (m *MockDataplaneSessionFactory) New(arg0 byte, arg1 int, arg2 addr.IA, arg3 net.Addr) control.DataplaneSession {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(control.DataplaneSession)
	return ret0
}

// New indicates an expected call of New
func (mr *MockDataplaneSessionFactoryMockRecorder) New(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockDataplaneSessionFactory)(nil).New), arg0, arg1, arg2, arg3)
}

// MockPktWriter is a mock of PktWriter interface
type MockPktWriter struct {
	ctrl     *gomock.Controller
	recorder *MockPktWriterMockRecorder
}

// MockPktWriterMockRecorder is the mock recorder for MockPktWriter
type MockPktWriterMockRecorder struct {
	mock *MockPktWriter
}

// NewMockPktWriter creates a new mock instance
func NewMockPktWriter(ctrl *gomock.Controller) *MockPktWriter {
	mock := &MockPktWriter{ctrl: ctrl}
	mock.recorder = &MockPktWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPktWriter) EXPECT() *MockPktWriterMockRecorder {
	return m.recorder
}

// Write mocks base method
func (m *MockPktWriter) Write(arg0 gopacket.Packet) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Write", arg0)
}

// Write indicates an expected call of Write
func (mr *MockPktWriterMockRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockPktWriter)(nil).Write), arg0)
}

// MockWorker is a mock of Worker interface
type MockWorker struct {
	ctrl     *gomock.Controller
	recorder *MockWorkerMockRecorder
}

// MockWorkerMockRecorder is the mock recorder for MockWorker
type MockWorkerMockRecorder struct {
	mock *MockWorker
}

// NewMockWorker creates a new mock instance
func NewMockWorker(ctrl *gomock.Controller) *MockWorker {
	mock := &MockWorker{ctrl: ctrl}
	mock.recorder = &MockWorkerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWorker) EXPECT() *MockWorkerMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockWorker) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockWorkerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockWorker)(nil).Close))
}

// Run mocks base method
func (m *MockWorker) Run() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run")
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockWorkerMockRecorder) Run() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockWorker)(nil).Run))
}

// MockSessionPolicyParser is a mock of SessionPolicyParser interface
type MockSessionPolicyParser struct {
	ctrl     *gomock.Controller
	recorder *MockSessionPolicyParserMockRecorder
}

// MockSessionPolicyParserMockRecorder is the mock recorder for MockSessionPolicyParser
type MockSessionPolicyParserMockRecorder struct {
	mock *MockSessionPolicyParser
}

// NewMockSessionPolicyParser creates a new mock instance
func NewMockSessionPolicyParser(ctrl *gomock.Controller) *MockSessionPolicyParser {
	mock := &MockSessionPolicyParser{ctrl: ctrl}
	mock.recorder = &MockSessionPolicyParserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSessionPolicyParser) EXPECT() *MockSessionPolicyParserMockRecorder {
	return m.recorder
}

// Parse mocks base method
func (m *MockSessionPolicyParser) Parse(arg0 []byte) (control.SessionPolicies, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parse", arg0)
	ret0, _ := ret[0].(control.SessionPolicies)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parse indicates an expected call of Parse
func (mr *MockSessionPolicyParserMockRecorder) Parse(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse", reflect.TypeOf((*MockSessionPolicyParser)(nil).Parse), arg0)
}

// MockRoutingPolicyProvider is a mock of RoutingPolicyProvider interface
type MockRoutingPolicyProvider struct {
	ctrl     *gomock.Controller
	recorder *MockRoutingPolicyProviderMockRecorder
}

// MockRoutingPolicyProviderMockRecorder is the mock recorder for MockRoutingPolicyProvider
type MockRoutingPolicyProviderMockRecorder struct {
	mock *MockRoutingPolicyProvider
}

// NewMockRoutingPolicyProvider creates a new mock instance
func NewMockRoutingPolicyProvider(ctrl *gomock.Controller) *MockRoutingPolicyProvider {
	mock := &MockRoutingPolicyProvider{ctrl: ctrl}
	mock.recorder = &MockRoutingPolicyProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRoutingPolicyProvider) EXPECT() *MockRoutingPolicyProviderMockRecorder {
	return m.recorder
}

// RoutingPolicy mocks base method
func (m *MockRoutingPolicyProvider) RoutingPolicy() *routing.Policy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RoutingPolicy")
	ret0, _ := ret[0].(*routing.Policy)
	return ret0
}

// RoutingPolicy indicates an expected call of RoutingPolicy
func (mr *MockRoutingPolicyProviderMockRecorder) RoutingPolicy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoutingPolicy", reflect.TypeOf((*MockRoutingPolicyProvider)(nil).RoutingPolicy))
}

// MockRunner is a mock of Runner interface
type MockRunner struct {
	ctrl     *gomock.Controller
	recorder *MockRunnerMockRecorder
}

// MockRunnerMockRecorder is the mock recorder for MockRunner
type MockRunnerMockRecorder struct {
	mock *MockRunner
}

// NewMockRunner creates a new mock instance
func NewMockRunner(ctrl *gomock.Controller) *MockRunner {
	mock := &MockRunner{ctrl: ctrl}
	mock.recorder = &MockRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRunner) EXPECT() *MockRunnerMockRecorder {
	return m.recorder
}

// Run mocks base method
func (m *MockRunner) Run(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockRunnerMockRecorder) Run(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockRunner)(nil).Run), arg0)
}

// MockGatewayWatcherFactory is a mock of GatewayWatcherFactory interface
type MockGatewayWatcherFactory struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayWatcherFactoryMockRecorder
}

// MockGatewayWatcherFactoryMockRecorder is the mock recorder for MockGatewayWatcherFactory
type MockGatewayWatcherFactoryMockRecorder struct {
	mock *MockGatewayWatcherFactory
}

// NewMockGatewayWatcherFactory creates a new mock instance
func NewMockGatewayWatcherFactory(ctrl *gomock.Controller) *MockGatewayWatcherFactory {
	mock := &MockGatewayWatcherFactory{ctrl: ctrl}
	mock.recorder = &MockGatewayWatcherFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGatewayWatcherFactory) EXPECT() *MockGatewayWatcherFactoryMockRecorder {
	return m.recorder
}

// New mocks base method
func (m *MockGatewayWatcherFactory) New(arg0 addr.IA, arg1 control.GatewayWatcherMetrics) control.Runner {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", arg0, arg1)
	ret0, _ := ret[0].(control.Runner)
	return ret0
}

// New indicates an expected call of New
func (mr *MockGatewayWatcherFactoryMockRecorder) New(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockGatewayWatcherFactory)(nil).New), arg0, arg1)
}

// MockPublisher is a mock of Publisher interface
type MockPublisher struct {
	ctrl     *gomock.Controller
	recorder *MockPublisherMockRecorder
}

// MockPublisherMockRecorder is the mock recorder for MockPublisher
type MockPublisherMockRecorder struct {
	mock *MockPublisher
}

// NewMockPublisher creates a new mock instance
func NewMockPublisher(ctrl *gomock.Controller) *MockPublisher {
	mock := &MockPublisher{ctrl: ctrl}
	mock.recorder = &MockPublisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPublisher) EXPECT() *MockPublisherMockRecorder {
	return m.recorder
}

// AddRoute mocks base method
func (m *MockPublisher) AddRoute(arg0 control.Route) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddRoute", arg0)
}

// AddRoute indicates an expected call of AddRoute
func (mr *MockPublisherMockRecorder) AddRoute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRoute", reflect.TypeOf((*MockPublisher)(nil).AddRoute), arg0)
}

// Close mocks base method
func (m *MockPublisher) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockPublisherMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPublisher)(nil).Close))
}

// DeleteRoute mocks base method
func (m *MockPublisher) DeleteRoute(arg0 control.Route) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteRoute", arg0)
}

// DeleteRoute indicates an expected call of DeleteRoute
func (mr *MockPublisherMockRecorder) DeleteRoute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRoute", reflect.TypeOf((*MockPublisher)(nil).DeleteRoute), arg0)
}

// MockPublisherFactory is a mock of PublisherFactory interface
type MockPublisherFactory struct {
	ctrl     *gomock.Controller
	recorder *MockPublisherFactoryMockRecorder
}

// MockPublisherFactoryMockRecorder is the mock recorder for MockPublisherFactory
type MockPublisherFactoryMockRecorder struct {
	mock *MockPublisherFactory
}

// NewMockPublisherFactory creates a new mock instance
func NewMockPublisherFactory(ctrl *gomock.Controller) *MockPublisherFactory {
	mock := &MockPublisherFactory{ctrl: ctrl}
	mock.recorder = &MockPublisherFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPublisherFactory) EXPECT() *MockPublisherFactoryMockRecorder {
	return m.recorder
}

// NewPublisher mocks base method
func (m *MockPublisherFactory) NewPublisher() control.Publisher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewPublisher")
	ret0, _ := ret[0].(control.Publisher)
	return ret0
}

// NewPublisher indicates an expected call of NewPublisher
func (mr *MockPublisherFactoryMockRecorder) NewPublisher() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewPublisher", reflect.TypeOf((*MockPublisherFactory)(nil).NewPublisher))
}
