<script lang="ts">
    import {onMount} from "svelte";
    import Cropper from "svelte-easy-crop";
    import type {CropArea} from "svelte-easy-crop/types";
    import {ImageCropArea} from "$lib/models/classes/imageCropArea";

    export let img:string
    export let avatarForm:HTMLFormElement
    export let cropArea:ImageCropArea = new ImageCropArea()
    let aspect = 1
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

    const sendImage = () => {
        avatarForm.requestSubmit()
    }

    let updateCrop = (e:CustomEvent) => {
        let pix:CropArea = e.detail.pixels
        cropArea.x = pix.x
        cropArea.y = pix.y
        cropArea.width = pix.width
        cropArea.height = pix.height
    };
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
            <div class="h-60 w-60 relative mx-auto py-3">
                <Cropper
                        image={img}
                        bind:zoom={cropArea.zoom}
                        bind:aspect
                        on:cropcomplete={updateCrop}
                        restrictPosition={true}
                />
            </div>

            <div class="w-60 mx-auto py-3 flex justify-between">
                <button on:click={sendImage}><i class="fa-regular fa-floppy-disk"></i> Save</button>
                <a on:click={closeModal} on:keydown={closeModal} class="cursor-pointer pt-2"><i class="fa-solid fa-ban"></i> Cancel</a>
            </div>

        </div>
    </div>
</div>