<script lang="ts">
    import Spinner from "$lib/components/Spinner.svelte";
    import {socketStateStore} from "$lib/stores/socketState.store";
    import {socketRetryCount} from "$lib/stores/socketRetryCount.store";
    import {onMount} from "svelte";
    let toaster:HTMLElement

    onMount(() => {
        toaster = document.getElementById("toaster")
    })

    export const openToaster = () => {
        toaster.classList.remove("hidden")
        toaster.classList.add("z-40")
    }

    export const closeToaster = () => {
        toaster.classList.add("hidden")
        toaster.classList.remove("z-40")
    }
</script>

<div class="fixed inset-0 z-30 overflow-y-auto h-full w-full hidden" id="toaster" on:mouseup={closeToaster}>
    <div class="top-20 max-w-screen-sm">
        <div class="bg-warning rounded py-3 mt-2 shadow-lg shadow-gray-800 w-[75%] mx-auto">
            <p class="text-center">
                <slot />
            </p>
        </div>
    </div>
</div>