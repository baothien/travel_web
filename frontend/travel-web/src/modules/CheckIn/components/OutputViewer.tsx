import { BackgroundConfig } from "../helpers/backgroundHelper";
import useRenderingPipeline from "../hooks/useRenderingPipeline";
import { TFLite } from "../hooks/useTFLite";

import { useEffect } from "react";
import { SourcePlayback } from "../helpers/sourceHelper";

type OutputViewerProps = {
  sourcePlayback: SourcePlayback;
  backgroundConfig: BackgroundConfig;
  tflite: TFLite;
};

function OutputViewer(props: OutputViewerProps) {
  const { pipeline, backgroundImageRef, canvasRef } = useRenderingPipeline(
    props.sourcePlayback,
    props.backgroundConfig,
    props.tflite
  );

  useEffect(() => {
    if (pipeline) {
      pipeline.updatePostProcessingConfig({
        jointBilateralFilter: { sigmaSpace: 1, sigmaColor: 0.1 },
        coverage: [0.5, 0.75],
        lightWrapping: 0.3,
      });
    }
  });

  return (
    <div>
      {props.backgroundConfig.type === "image" && (
        <img
          ref={backgroundImageRef}
          src={props.backgroundConfig.url}
          alt=""
          hidden={true}
        />
      )}

      <canvas
        id="canvas-output"
        key="webgl2"
        ref={canvasRef}
        width={props.sourcePlayback.width}
        height={props.sourcePlayback.height}
      />
    </div>
  );
}

export default OutputViewer;
