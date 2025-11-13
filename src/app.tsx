import { TextAttributes } from "@opentui/core";
import { render, useKeyboard, useRenderer, useTerminalDimensions } from "@opentui/solid";
import { createEffect, createSignal, Match, onMount, Show, Switch } from "solid-js";

import Home from "./components/Home";
import Blog from "./components/Blog";
import Header from "./components/Header";
import Footer from "./components/Footer";

const TABS = [
  { name: "home", scene: "home", description: "home desc", keys: [{bind:"b", desc:"somet"}]},
  { name: "blog", scene: "blog", description: "blog desc", keys: [{bind: "2",desc: "blog key"}] },
  { name: "resume", scene: "resume", description: "resume desc", keys: [] },
  { name: "contact", scene: "contact", description: "contact desc", keys: [] },
];

export default function App() {
  const renderer = useRenderer();
  onMount(() => {
    renderer.useConsole = true;
  });
  const dimensions = useTerminalDimensions();
  const terminalBigEnough = () => (dimensions().width >= 80 && dimensions().height >= 30);

  const contentHeight = () => {
    const real = dimensions().height;
    if (real <= 30) return 24;
    return Math.min(real, 42) - 8;
  };
  const contentWidth = () => {
    const real = dimensions().width;
    if (real <= 80) return 80;
    if (real <= 100) return real;
    return Math.min(real, 120);
  };

  createEffect(() => {
    console.log("contentWidth:", contentWidth(), "contentHeight:", contentHeight(), "realWidth:", dimensions().width, "realHeight:", dimensions().height);
  });

  const [activeTab, setActiveTab] = createSignal(0);

  useKeyboard((key) => {
    switch (key.name) {
      case "`":
        renderer.console.toggle();
        break;
      case "h":
        setActiveTab(0)
        break;
      case "b":
        setActiveTab(1);
        break;
      case "r":
        setActiveTab(2);
        break;
      case "c":
        setActiveTab(3);
        break;
      case "left":
        setActiveTab((prev) => (Math.max(prev - 1, 0)));
        break;
      case "right":
        setActiveTab((prev) => (( prev + 1) % TABS.length));
        break;
      case "t":
        renderer.toggleDebugOverlay();
        break;
    }
  });

  const SizeFallback = () => {
    return (
      <box alignItems="center">
        <text marginBottom={1}>Your terminal is too small!</text>
        <text attributes={TextAttributes.DIM}>{dimensions().width}x{dimensions().height} {"<"} 80x30</text>
      </box>
    );
  };

  return (
    <box flexDirection="column" alignItems="center" justifyContent="center" height="100%" width="100%">
      <Show when={terminalBigEnough()} fallback={<SizeFallback />}>
        <Header width={contentWidth()} options={TABS} activeTab={activeTab} />
        <box height={contentHeight()} width={contentWidth() - 2}>
          <Switch>
            <Match when={activeTab() === 0}>
              <Home />
            </Match>
            <Match when={activeTab() == 1}>
              <Blog />
            </Match>
            <Match when={activeTab() == 2}>
              <text>Resume</text>
            </Match>
            <Match when={activeTab() == 3}>
              <text>Contact</text>
            </Match>
          </Switch>
        </box>
        <Footer width={contentWidth()} keys={() => (TABS.at(activeTab())?.keys)} />
      </Show>
    </box>
  );
}

if (import.meta.main) {
  render(App, {
    useKittyKeyboard: false,
    consoleOptions: {
      maxStoredLogs: 1000,
      sizePercent: 40,
    },
  })
}
