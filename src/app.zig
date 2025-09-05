const std = @import("std");
const vaxis = @import("vaxis");
const vxfw = vaxis.vxfw;
const Key = vaxis.Key;

const texts = @import("text.zig");

const Bind = struct {
    key: vaxis.Key,
    action: Action,

    const Action = union(enum) {
        quit,
        redraw,
        @"next-tab",
        @"prev-tab",
        @"select-tab": u8,
        @"scroll-up",
        @"scroll-down",
    };
};

pub const App = struct {
    const minimum_size: vxfw.Size = .{ .width = 80, .height = 30 };

    alloc: std.mem.Allocator,
    // tabs: [4][]const u8,
    page: enum { home, blog },

    binds: std.ArrayList(Bind),

    pub fn init(self: *App, gpa: std.mem.Allocator) !void {
        self.* = .{
            .alloc = gpa,
            .binds = try .initCapacity(gpa, 16),
            .page = .home,
        };
        errdefer self.deinit();

        try self.binds.append(.{ .key = .{ .codepoint = 'q' }, .action = .quit });
        try self.binds.append(.{ .key = .{ .codepoint = 'j' }, .action = .@"scroll-down" });
        try self.binds.append(.{ .key = .{ .codepoint = 'k' }, .action = .@"scroll-up" });
        // try self.binds.append(.{ .key = .{ .codepoint = 'h' }, .action = .@"prev-tab" });
        try self.binds.append(.{ .key = .{ .codepoint = 'l' }, .action = .@"next-tab" });
        try self.binds.append(.{ .key = .{ .codepoint = Key.down }, .action = .@"scroll-down" });
        try self.binds.append(.{ .key = .{ .codepoint = Key.up }, .action = .@"scroll-up" });
        try self.binds.append(.{ .key = .{ .codepoint = Key.left }, .action = .@"prev-tab" });
        try self.binds.append(.{ .key = .{ .codepoint = Key.right }, .action = .@"next-tab" });
        try self.binds.append(.{ .key = .{
            .codepoint = 'l',
            .mods = .{ .ctrl = true },
        }, .action = .redraw });
    }

    pub fn deinit(self: *App) void {
        self.binds.deinit();
    }

    pub fn widget(self: *App) vxfw.Widget {
        return .{
            .userdata = self,
            .eventHandler = App.typeErasedEventHandler,
            .drawFn = App.typeErasedDrawFn,
        };
    }

    /// This function will be called from the vxfw runtime.
    fn typeErasedEventHandler(ptr: *anyopaque, ctx: *vxfw.EventContext, event: vxfw.Event) anyerror!void {
        const self: *App = @ptrCast(@alignCast(ptr));
        switch (event) {
            .init => return ctx.requestFocus(self.widget()),
            .key_press => |key| {
                if (key.matches('c', .{ .ctrl = true })) {
                    ctx.quit = true;
                    return;
                }
                for (self.binds.items) |bind| {
                    if (key.matches(bind.key.codepoint, bind.key.mods)) {
                        switch (bind.action) {
                            .quit => ctx.quit = true,
                            // .@"next-tab" => self.nextTab(),
                            // .@"prev-tab" => self.prevTab(),
                            // .@"select-tab" => |idx| self.selectTab(idx),
                            .redraw => try ctx.queueRefresh(),
                            else => {},
                        }
                        return ctx.consumeAndRedraw();
                    }
                }

                if (key.matches('h', .{})) {
                    self.page = .home;
                    ctx.consumeAndRedraw();
                    return;
                }

                if (key.matches('b', .{})) {
                    self.page = .blog;
                    ctx.consumeAndRedraw();
                    return;
                }
            },
            // We can request a specific widget gets focus. In this case, we always want to focus
            // our button. Having focus means that key events will be sent up the widget tree to
            // the focused widget, and then bubble back down the tree to the root. Users can tell
            // the runtime the event was handled and the capture or bubble phase will stop
            // .focus_in => return ctx.requestFocus(self.button.widget()),
            else => {},
        }
    }

    /// This function is called from the vxfw runtime. It will be called on a regular interval, and
    /// only when any event handler has marked the redraw flag in EventContext as true. By
    /// explicitly requiring setting the redraw flag, vxfw can prevent excessive redraws for events
    /// which don't change state (ie mouse motion, unhandled key events, etc)
    fn typeErasedDrawFn(ptr: *anyopaque, ctx: vxfw.DrawContext) std.mem.Allocator.Error!vxfw.Surface {
        const self: *App = @ptrCast(@alignCast(ptr));
        const max_size = ctx.max.size();
        if (max_size.width < minimum_size.width or max_size.height < minimum_size.height) {
            return self.drawResizeView(ctx);
        }

        var children = std.ArrayList(vxfw.SubSurface).init(ctx.arena);

        const header_text: vxfw.Text = .{ .text = "david_parrott" };
        const header: vxfw.Border = .{ .child = header_text.widget() };
        const header_child: vxfw.SubSurface = .{
            .origin = .{ .row = 0, .col = 0 },
            .surface = try header.draw(ctx),
        };
        try children.append(header_child);

        const flex_items = try ctx.arena.alloc(vxfw.FlexItem, 2);
        const home: vxfw.Text = .{ .text = "home", .style = .{ .fg = .{ .index = if (self.page == .home) 2 else 7 } } };
        const blog: vxfw.Text = .{ .text = "blog", .style = .{ .fg = .{ .index = if (self.page == .blog) 2 else 7 } } };
        flex_items[0] = vxfw.FlexItem.init(home.widget(), 1);
        flex_items[1] = vxfw.FlexItem.init(blog.widget(), 1);

        const flex_row: vxfw.FlexRow = .{ .children = flex_items };
        const flex_border: vxfw.Border = .{ .child = flex_row.widget() };
        try children.append(.{
            .origin = .{ .row = 4, .col = 0 },
            .surface = try flex_border.draw(ctx),
        });

        var center: vxfw.Center = undefined;

        return switch (self.page) {
            .home => blk: {
                const bird_render: vxfw.Text = .{
                    .text = texts.bird,
                };
                center.child = bird_render.widget();
                try children.append(.{
                    .origin = .{ .row = 7, .col = 0 },
                    .surface = try center.draw(ctx),
                });
                break :blk .{
                    .size = max_size,
                    .widget = self.widget(),
                    .buffer = &.{},
                    .children = children.items,
                };
            },
            .blog => blk: {
                const blog_render: vxfw.Text = .{
                    .text = texts.about.blog,
                };
                center.child = blog_render.widget();
                try children.append(.{
                    .origin = .{ .row = 7, .col = 0 },
                    .surface = try center.draw(ctx),
                });
                break :blk .{
                    .size = max_size,
                    .widget = self.widget(),
                    .buffer = &.{},
                    .children = children.items,
                };
            },
        };
    }

    fn drawResizeView(self: *App, ctx: vxfw.DrawContext) std.mem.Allocator.Error!vxfw.Surface {
        const fmt = std.fmt.comptimePrint(
            "your terminal is too small ({{d}}x{{d}})\nresize to >= {d}x{d}",
            .{ minimum_size.width, minimum_size.height },
        );

        const term_size = ctx.max.size();
        const resize_msg = try std.fmt.allocPrint(ctx.arena, fmt, .{ term_size.width, term_size.height });
        var text = try ctx.arena.create(vxfw.Text);
        text.* = .{
            .text = resize_msg,
            .text_align = .center,
        };

        var boxed = try ctx.arena.create(vxfw.Center);
        boxed.* = .{ .child = text.widget() };
        const subsurface = try ctx.arena.alloc(vxfw.SubSurface, 1);
        subsurface[0] = .{
            .origin = .{ .row = 0, .col = 0 },
            .surface = try boxed.draw(ctx),
        };

        return .{
            .size = term_size,
            .widget = self.widget(),
            .buffer = &.{},
            .children = subsurface,
        };
    }
};
