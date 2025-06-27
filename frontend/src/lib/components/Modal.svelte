<script lang="ts">
    import CloseCircleOutline from "svelte-material-icons/CloseCircleOutline.svelte";
    import { fade } from 'svelte/transition';
    import {cubicInOut} from "svelte/easing";

    
    interface Props {
        // this flag makes it so the modal will close if the user clicks anywhere on the screen
        isEasilyClosable?: boolean;
        children?: import('svelte').Snippet;
        openModal?: () => void;
        closeModal?: () => void;
    }

    let { isEasilyClosable = false, children, openModal = $bindable(), closeModal = $bindable() }: Props = $props();
    let shouldDisplay = $state(false)

    openModal = () => {
        shouldDisplay = true

    }

    closeModal = () => {
        shouldDisplay = false

    }
    const closeModalCheck = () => {
        if (isEasilyClosable) {
            closeModal()
        }

    }
    let optionalPadding = $derived(isEasilyClosable ? "" : "p-3")
</script>

{#if shouldDisplay}
    <!-- Overlay -->
    <div transition:fade|global={{ duration: 300, easing: cubicInOut }} class="fixed inset-0 z-10 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full" onmouseup={closeModalCheck} role="button" tabindex="0">

        <!-- Modal window -->
        <div class="relative top-20 {optionalPadding} border border-secondary max-w-screen-sm mx-auto shadow-lg rounded-md bg-canvas-secondary">

            <div>

                {#if !isEasilyClosable}
                    <!-- Close button -->
                    <div class="absolute top-2 right-3">
                        <button class="bg-transparent" onclick={closeModal}>
                            <CloseCircleOutline size="1.5em"/>
                        </button>
                    </div>
                {/if}

                <!-- Content -->
                {@render children?.()}

            </div>
        </div>
    </div>
{/if}