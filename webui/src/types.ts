export interface DiveDesc {
    id: number
    label: string
    cardinal: number
}

export interface DiveTripDesc {
    id: number
    trip: string
    descriptors: DiveDesc[]
}
