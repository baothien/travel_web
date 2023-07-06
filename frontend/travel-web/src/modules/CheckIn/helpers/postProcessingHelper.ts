export type BlendMode = 'screen' | 'linearDodge'

export type PostProcessingConfig = {
  // smoothSegmentationMask: boolean
  jointBilateralFilter: JointBilateralFilterConfig
  coverage: [number, number]
  lightWrapping: number
}

export type JointBilateralFilterConfig = {
  sigmaSpace: number
  sigmaColor: number
}
