export interface DiveDesc {
    id: number
    label: string
    cardinal: number
}

export interface DiveTripDesc {
    id: number
    label: string
    descriptors: DiveDesc[]
}

export interface DiveSiteDesc {
    id: number
    name: string
}

export interface Dive {
    id: number
    label: string
    cardinal: number
    max_depth: string
}

export interface DiveSite {
    id: number
    name: string
    coordinates: string
}

export interface StatusResponse {
    ok: boolean
    status: number | null
    error: any
}
