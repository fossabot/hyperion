import React, { Component } from "react";
import { connect } from "react-redux";
import { fetchCueMaster } from "../actions/cues";
import {
  CueTable,
  CueTableCol,
  CueLabel,
  CueFrame,
  CueFrameWait,
  CueFrameWrapper,
  Progress
} from "../components/Cue";
import { bindActionCreators } from "redux";
import { Header } from "semantic-ui-react";
class cueList extends Component {
  componentDidMount() {
    this.props.fetchCueMaster();
  }

  render() {
    let mainStack = this.props.stacks[0];

    if (!mainStack) return <div>loading...</div>;

    const bare = { length_ms: 0, items: [] };

    let all = {};
    let cuesById = {};

    let cueList = mainStack.processed_cues.concat(
      mainStack.cues || [],
      mainStack.active_cue || []
    );

    cueList.sort((a, b) => a.id - b.id);

    cueList.forEach(c => {
      let maxActions = 1;
      c.frames.forEach(
        f => (maxActions = Math.max(maxActions, f.actions.length))
      );
      all[c.id] = Array.apply(null, Array(maxActions)).map(x => bare);
      cuesById[c.id] = c;
    });

    cueList.forEach(c =>
      c.frames.forEach((f, z) => {
        f.actions.forEach((action, x) => {
          let tmp = {};
          Object.assign(tmp, all[c.id][x]);
          tmp["length_ms"] += action.action_duration_ms;
          tmp["items"] = tmp["items"].slice();
          tmp["items"].push(
            <CueFrame
              action={action}
              key={c.id + "-" + z + "-" + x}
              duration={action.action_duration_ms}
              frameId={f.id}
              actionId={action.id}
            />
          );
          all[c.id][x] = tmp;
        });
        //todo: add padding
        let targetLen = all[c.id][0]["length_ms"];
        all[c.id].forEach(
          x => (targetLen = Math.max(targetLen, x["length_ms"]))
        );
        all[c.id].forEach((item, x) => {
          let padding = targetLen - item["length_ms"];
          if (padding !== 0) {
            let tmp = {};
            Object.assign(tmp, all[c.id][x]);
            tmp["length_ms"] += padding;
            tmp["items"].push(<CueFrameWait key={x} duration={padding} />);
            all[c.id][x] = tmp;
          }
        });
      })
    );

    return (
      <div>
        <Header content={"cues"} />
        <CueTable>
          <CueTableCol>
            {cueList.map(c => {
              let maxActions = 1;
              c.frames.forEach(
                f => (maxActions = Math.max(maxActions, f.actions.length))
              );
              return (
                <CueLabel
                  id={c.id}
                  key={c.id}
                  numActions={maxActions}
                  status={c.status}
                  cue={c}
                  duration={c.expected_duration_ms}
                  duration_drift_ms={c.duration_drift_ms}
                />
              );
            })}
          </CueTableCol>
          <CueTableCol>
            {Object.keys(all).map(k => {
              return (
                <div>
                  <CueFrameWrapper key={k + "-2"}>
                    <Progress cue={cuesById[parseInt(k, 10)]} />
                  </CueFrameWrapper>
                  {all[k].map((item, x) => (
                    <CueFrameWrapper key={x}> {item.items}</CueFrameWrapper>
                  ))}
                </div>
              );
            })}
          </CueTableCol>
        </CueTable>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return { stacks: state.cues.cue_stacks };
}

const mapDispatchToProps = dispatch => {
  return bindActionCreators({ fetchCueMaster }, dispatch);
};

export default connect(mapStateToProps, mapDispatchToProps)(cueList);
