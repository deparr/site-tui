import { For, type Setter, type Accessor } from "solid-js";
import { colors } from "../lib/theme";

export default function Header(props: {
  options: any[],
  width: number
  activeTab: Accessor<number>
}) {
  return (
    <box border borderColor={colors.dark.border} flexDirection="row" width={props.width} justifyContent="space-evenly">
      <box flexDirection="row">
        <text fg={colors.dark.accent}><em><strong>{"@"}</strong></em></text>
        <text fg={colors.dark.highlight}><em><strong>dp</strong></em></text>
      </box>
      <text fg={colors.dark.border}>│</text>
      <For each={props.options} >
        {(item, index) => (
          <>
            <box flexDirection="row">
              <text fg={index() == props.activeTab() ? colors.dark.accent : colors.dark.border}
              ><u>{item.name[0]}</u>{item.name.substr(1)}</text>
            </box>
            {index() < props.options.length - 1 ? <text fg={colors.dark.border}>│</text> : undefined}
          </>
        )}
      </For>
    </box>
  );
}
