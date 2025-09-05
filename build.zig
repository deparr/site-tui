const std = @import("std");

pub fn build(b: *std.Build) void {
    const target = b.standardTargetOptions(.{});
    const optimize = b.standardOptimizeOption(.{});

    const exe_mod = b.createModule(.{
        .root_source_file = b.path("src/main.zig"),
        .target = target,
        .optimize = optimize,
    });

    const vaxis = b.dependency("vaxis", .{
        .target =  target,
        .optimize = optimize,
    });

    exe_mod.addImport("vaxis", vaxis.module("vaxis"));

    const exe = b.addExecutable(.{
        .name = "tui",
        .root_module = exe_mod,
    });

    const exe_check = b.addExecutable(.{
        .name = "tui",
        .root_module = exe_mod,
    });

    if (b.option(bool, "no-bin", "no emit bin") orelse false) {
        b.getInstallStep().dependOn(&exe_check.step);
    } else {
        b.installArtifact(exe);
    }
}
