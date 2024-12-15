import type { Dive } from "@/types"

export function sleep(ms: number): Promise<void> {
    return new Promise(resolve => setTimeout(resolve, ms))
}

export function formatJSON(json: object): string {
    return JSON.stringify(json, null, 4)
}

export function saveStringToLocalStorage(key: string, value: string): void {
    localStorage.setItem(key, value)
}

export function loadStringFromLocalStorage(key: string): string | null {
    return localStorage.getItem(key)
}

export function diveIdToRoute(id: number): string {
    return `/dives/${paddedID(id)}`
}

export function siteIdToRoute(id: number): string {
    return `/sites/${paddedID(id)}`
}

export function paddedID(id: number): string {
    return id.toString().padStart(4, '0')
}

export function composeDiveTitle(dive: Dive): string {
    return `Dive ${dive.cardinal}. ${dive.label}`
}
