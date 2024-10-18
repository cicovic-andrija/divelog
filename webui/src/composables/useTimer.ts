import { ref, computed } from 'vue'

export function useTimer() {
    const seconds = ref(0)
    const TIMEOUT: number = 10
    const timedOut = computed(() => seconds.value > TIMEOUT)

    let interval: number = 0
    function start() {
        interval = setInterval(() => {
            if (seconds.value <= TIMEOUT) {
                seconds.value++
            }
        }, 1000)
    }

    function stop() {
        clearInterval(interval)
    }

    return { seconds, timedOut, start, stop }
}
