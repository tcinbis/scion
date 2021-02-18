nogo_deps = [
    "//go/lint:log",
    "@com_github_oncilla_gochecks//logcheck:go_tool_library",
    "@com_github_oncilla_gochecks//serrorscheck:go_tool_library",
    "@org_golang_x_tools//go/analysis/passes/asmdecl:go_tool_library",
    "@org_golang_x_tools//go/analysis/passes/assign:go_tool_library",
    "@org_golang_x_tools//go/analysis/passes/atomic:go_tool_library",
    "@org_golang_x_tools//go/analysis/passes/atomicalign:go_tool_library",
    "@org_golang_x_tools//go/analysis/passes/bools:go_tool_library",
    "@org_golang_x_tools//go/analysis/passes/buildtag:go_tool_library",
    # This crashes the build of @com_github_vishvananda_netlink
    # "@org_golang_x_tools//go/analysis/passes/cgocall:go_tool_library",
    "@org_golang_x_tools//go/analysis/passes/composite:go_tool_library",
    "@org_golang_x_tools//go/analysis/passes/copylock:go_tool_library",
    "@org_golang_x_tools//go/analysis/passes/deepequalerrors:go_tool_library",
    "@org_golang_x_tools//go/analysis/passes/errorsas:go_tool_library",
    # We have uncountable violations of fieldalignment.
    # "@org_golang_x_tools//go/analysis/passes/fieldalignment:go_tool_library",
    "@org_golang_x_tools//go/analysis/passes/httpresponse:go_tool_library",
    "@org_golang_x_tools//go/analysis/passes/ifaceassert:go_tool_library",
    "@org_golang_x_tools//go/analysis/passes/loopclosure:go_tool_library",
    "@org_golang_x_tools//go/analysis/passes/lostcancel:go_tool_library",
    "@org_golang_x_tools//go/analysis/passes/nilfunc:go_tool_library",
    "@org_golang_x_tools//go/analysis/passes/nilness:go_tool_library",
    "@org_golang_x_tools//go/analysis/passes/printf:go_tool_library",
    "@org_golang_x_tools//go/analysis/passes/shift:go_tool_library",
    "@org_golang_x_tools//go/analysis/passes/sortslice:go_tool_library",
    "@org_golang_x_tools//go/analysis/passes/stdmethods:go_tool_library",
    "@org_golang_x_tools//go/analysis/passes/stringintconv:go_tool_library",
    "@org_golang_x_tools//go/analysis/passes/structtag:go_tool_library",
    "@org_golang_x_tools//go/analysis/passes/tests:go_tool_library",
    "@org_golang_x_tools//go/analysis/passes/unmarshal:go_tool_library",
    "@org_golang_x_tools//go/analysis/passes/unreachable:go_tool_library",
    "@org_golang_x_tools//go/analysis/passes/unsafeptr:go_tool_library",
    "@org_golang_x_tools//go/analysis/passes/unusedresult:go_tool_library",
]
