<script lang="ts">
    import {onMount} from "svelte";
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
</script>

<!-- Overlay -->
<div class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full hidden" id="modal">

    <!-- Modal window -->
    <div class="relative top-20 p-3 border border-secondary max-w-screen-sm mx-auto shadow-lg rounded-md bg-canvas-secondary">

        <div>
            <!-- Close button -->
            <div class="absolute top-2 right-3">
                <button class="bg-transparent" on:click={closeModal}><i class="fa-regular fa-circle-xmark"></i></button>
            </div>

            <!-- Content -->
            <slot />

        </div>
    </div>
</div>