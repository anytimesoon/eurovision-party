<script lang="ts">
    import {onMount} from "svelte";
    import CloseCircleOutline from "svelte-material-icons/CloseCircleOutline.svelte";

    // this flag makes it so the modal will close if the user clicks anywhere on the screen
    export let isEasilyClosable:boolean = false
    let modal:HTMLElement

    onMount(() => {
        modal = document.getElementById("modal")
    })

    export const openModal = () => {
        modal.classList.remove("hidden")
        modal.classList.add("z-50")
    }

    export const closeModal = () => {
        modal.classList.add("hidden")
        modal.classList.remove("z-50")
    }

    const closeModalCheck = () => {
        if (isEasilyClosable) {
            closeModal()
        }
    }

    $: optionalPadding = isEasilyClosable ? "" : "p-3"
</script>

<!-- Overlay -->
<div class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full hidden" id="modal" on:mouseup={closeModalCheck}>

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