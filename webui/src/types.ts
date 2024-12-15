export interface DiveSiteDesc {
    id: number
    name: string
}

export interface DiveSite {
    id: number
    name: string
    coordinates: string
    geo_labels: string[]
    linked_dives: {
        id: number
        short_label: string
        date_time: string
    }
}

export interface DiveDesc {
    id: number
    short_label: string
    date_time: string
}

export interface DiveTripDesc {
    id: number
    label: string
    linked_dives: DiveDesc[]
}

export interface Dive {
    id: number
    label: string
    cardinal: number
    max_depth: string
}

export interface StatusResponse {
    ok: boolean
    status: number | null
    error: any
}
