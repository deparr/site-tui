const std = @import("std");
const vaxis = @import("vaxis");
const vxfw = vaxis.vxfw;

const App = @import("app.zig").App;

pub fn main() !void {
    var gpa: std.heap.DebugAllocator(.{}) = .init;
    defer _ = gpa.deinit();

    const allocator = gpa.allocator();

    var app = try vxfw.App.init(allocator);
    defer app.deinit();

    const site_app = try allocator.create(App);
    defer allocator.destroy(site_app);
    defer site_app.deinit();

    try site_app.init(allocator);

    try app.run(site_app.widget(), .{});
}
