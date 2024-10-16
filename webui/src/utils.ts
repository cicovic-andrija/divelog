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
