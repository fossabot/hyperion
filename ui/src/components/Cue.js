import React from "react";
import styled from "styled-components";
import { rgbToHex } from "../utils";

const COLOR_PROCESSED = "#B360E4";
const COLOR_ENQUEUED = "#008AFF";
const COLOR_ACTIVE = "#56D868";
export const CueFrameWrapper = styled.div`
  display: flex;
`;

export const CueTable = styled.div`
  display: flex;
  border: 1px solid black;
`;
export const CueTableCol = styled.div`
  flex-direction: column;
`;

const TIME_SCALE = 0.3;

const ColorPreview = styled.div.attrs({
  style: ({ hex }) => ({
    backgroundColor: hex
  })
})`
  min-width: 15px;
  width: 15px;
  min-height: 15px;
  height: 15px;
`;

const CueFrameInner = styled.div.attrs({
  style: ({ duration }) => ({
    width: `${duration * TIME_SCALE}px`,
    minWidth: `${duration * TIME_SCALE}px`
  })
})`
  height: 70px;
  border: 1px solid black;
  padding: 5px;
`;

export const CueFrame = ({ ...props }) => {
  let { rgb } = props.action.new_state;
  let hex = rgbToHex(rgb.r, rgb.g, rgb.b);
  return (
    <CueFrameInner {...props}>
      {props.duration} ms (F:{props.frameId} | A:{props.actionId}) <br />
      {props.action.light_name} <ColorPreview hex={hex} />
    </CueFrameInner>
  );
};

const CueFrameWaitInner = styled.div.attrs({
  style: ({ duration }) => ({
    width: `${duration * TIME_SCALE}px`,
    minWidth: `${duration * TIME_SCALE}px`
  })
})`
  height: 70px;
  border: 1px solid black;
  background-color: #f96f3a;
  padding: 5px;
`;
export const CueFrameWait = ({ ...props }) => (
  <CueFrameWaitInner {...props}> {props.duration} ms wait </CueFrameWaitInner>
);

const CueLabelInner = styled.div.attrs({
  style: ({ numActions, status }) => {
    let statusColor = COLOR_ENQUEUED;
    if (status === "active") statusColor = COLOR_ACTIVE;
    else if (status === "processed") statusColor = COLOR_PROCESSED;
    return {
      height: `${numActions * 70 + 20}px`,
      backgroundColor: statusColor
    };
  }
})`
  width: 100px;
  border: 1px solid black;
  color: white;
  display: flex;
  justify-content: center;
  flex-direction: column;
  text-align: center;
  order: 0;
`;

export const CueLabel = ({ ...props }) => {
  let { id, duration, duration_drift_ms, cue } = props;
  return (
    <CueLabelInner {...props}>
      # {id} <br />
      {`${duration} ms`}{" "}
      {cue.status === "active"
        ? `${(cue.elapsed_ms / cue.expected_duration_ms * 100).toFixed(1)} %`
        : null}
      <i>{(duration_drift_ms && `(+${duration_drift_ms} ms)`) || null}</i>
    </CueLabelInner>
  );
};

const ProgressInner = styled.div.attrs({
  style: ({ duration, color }) => ({
    width: `${duration * TIME_SCALE}px`,
    minWidth: `${duration * TIME_SCALE}px`,
    backgroundColor: color
  })
})`
  height: 20px;
  border: 1px solid black;
`;
export const Progress = ({ ...props }) => {
  let { cue } = props;
  let { elapsed_ms, expected_duration_ms, status } = cue;
  let duration = 0;
  let color = COLOR_PROCESSED;
  if (status === "active") {
    duration = Math.min(elapsed_ms, expected_duration_ms);
    color = COLOR_ACTIVE;
  }
  if (status === "processed") duration = expected_duration_ms;

  return <ProgressInner duration={duration} color={color} />;
};
