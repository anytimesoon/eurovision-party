<script lang="ts">
    import CloseCircleOutline from "svelte-material-icons/CloseCircleOutline.svelte";
    import { fade } from 'svelte/transition';
    import {cubicInOut} from "svelte/easing";

    // this flag makes it so the modal will close if the user clicks anywhere on the screen
    export let isEasilyClosable:boolean = false
    let shouldDisplay = false

    export const openModal = () => {
        console.log("hello")
        shouldDisplay = true
    }

    export const closeModal = () => {
        shouldDisplay = false
    }

    const closeModalCheck = () => {
        if (isEasilyClosable) {
            closeModal()
        }
    }

    $: optionalPadding = isEasilyClosable ? "" : "p-3"
</script>

{#if shouldDisplay}
    <!-- Overlay -->
    <div transition:fade={{ duration: 300, easing: cubicInOut }} class="fixed inset-0 z-10 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full" on:mouseup={closeModalCheck}>

        <!-- Modal window -->
        <div class="relative top-20 {optionalPadding} border border-secondary max-w-screen-sm mx-auto shadow-lg rounded-md bg-canvas-secondary">

            <div>

                {#if !isEasilyClosable}
                    <!-- Close button -->
                    <div class="absolute top-2 right-3">
                        <button class="bg-transparent" on:click={closeModal}>
                            <CloseCircleOutline size="1.5em"/>
                        </button>
                    </div>
                {/if}

                <!-- Content -->
                <slot />

            </div>
        </div>
    </div>
{/if}