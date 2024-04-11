<script lang="ts">
    import { fade } from 'svelte/transition';

    let shouldDisplay = false

    export const openToaster = () => {
        shouldDisplay = true
        setTimeout(closeToaster, 2000)
    }

    export const closeToaster = () => {
        shouldDisplay = false
    }
</script>

{#if shouldDisplay}
    <div transition:fade={{ delay: 250, duration: 300 }} class="fixed inset-0 z-30 overflow-y-auto h-full w-full z-40" on:mouseup={closeToaster}>
        <div class="top-20 max-w-screen-sm">
            <div class="bg-warning rounded py-3 px-1.5 mt-2 shadow-lg shadow-gray-800 w-[75%] mx-auto">
                <p class="text-center">
                    <slot />
                </p>
            </div>
        </div>
    </div>
{/if}