// import { BodyPix } from '@tensorflow-models/body-pix'
import { useEffect, useRef, useState } from 'react'
import { BackgroundConfig } from '../helpers/backgroundHelper'
import { RenderingPipeline } from '../helpers/renderingPipelineHelper'
import { SourcePlayback } from '../helpers/sourceHelper'
import { createTimerWorker } from '../helpers/timerHelper'
import { buildWebGL2Pipeline } from '../pipelines/webgl2/webgl2Pipeline'
import { TFLite } from './useTFLite'

function useRenderingPipeline(
  sourcePlayback: SourcePlayback,
  backgroundConfig: BackgroundConfig,
  tflite: TFLite
) {
  const [pipeline, setPipeline] = useState<RenderingPipeline | null>(null)
  const backgroundImageRef = useRef<HTMLImageElement>(null)
  const canvasRef = useRef<HTMLCanvasElement>(null!)

  useEffect(() => {
    const targetTimerTimeoutMs = 1000 / 65

    let previousTime = 0
    let beginTime = 0
    let eventCount = 0
    let frameCount = 0
    const frameDurations: number[] = []

    let renderTimeoutId: number

    const timerWorker = createTimerWorker()

    const newPipeline = buildWebGL2Pipeline(
      sourcePlayback,
      backgroundImageRef.current,
      backgroundConfig,
      canvasRef.current,
      tflite,
      addFrameEvent
    )

    async function render() {
      const startTime = performance.now()

      beginFrame()
      await newPipeline.render()
      endFrame()

      renderTimeoutId = timerWorker.setTimeout(
        render,
        Math.max(0, targetTimerTimeoutMs - (performance.now() - startTime))
      )
    }

    function beginFrame() {
      beginTime = Date.now()
    }

    function addFrameEvent() {
      const time = Date.now()
      frameDurations[eventCount] = time - beginTime
      beginTime = time
      eventCount++
    }

    function endFrame() {
      const time = Date.now()
      frameDurations[eventCount] = time - beginTime
      frameCount++
      if (time >= previousTime + 1000) {
        previousTime = time
        frameCount = 0
      }
      eventCount = 0
    }

    render()

    console.log('Animation started:', sourcePlayback, backgroundConfig)

    setPipeline(newPipeline)

    return () => {
      timerWorker.clearTimeout(renderTimeoutId)
      timerWorker.terminate()
      newPipeline.cleanUp()
      console.log('Animation stopped:', sourcePlayback, backgroundConfig)

      setPipeline(null)
    }
  }, [sourcePlayback, backgroundConfig, tflite])

  return {
    pipeline,
    backgroundImageRef,
    canvasRef,
  }
}

export default useRenderingPipeline
