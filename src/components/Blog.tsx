import { tui } from "../lib/theme";

const codeContent = `//! toplevel doc

const std = @import("std");

/// which direction a entity can face
const Dir = enum {
    north,
    east,
    south,
    west,
};

const Entity = struct {
    pos: struct { x: f32, y: f32 },
    id: usize = 0,
    dir: Dir = .north,

    pub fn dist_to(self: *const Entity, other: *const Entity) f32 {
        const yy = other.pos.y - self.pos.y;
        const xx = other.pos.x - self.pos.x;
        return @sqrt(yy * yy + xx * xx)
    }
}

pub fn main() !void {
    const stdout = std.io.getStdOut().writer();
    try stdout.print("Hello, {}!\n", .{"world"});
    const current_dir = .east;
    var day = switch(current_day) {
        .north => "North",
        .east => "East",
        .south => "South",
        .west => "West",
    }
    var maybe_direction: ?Dir = null;
    const target = std.Target.current.os.tag;
    var entity_pool: []*Entity = undefined;
    var c = 'c';
    var x = 0xc0ffee;
    var a = @fieldParentPtr("somefield", stdout);
}
`;

export default function Blog() {
  return (<code content={codeContent} filetype="zig" syntaxStyle={tui.syntax.dark} />);
}
