<script lang="ts">
    import {currentUser} from "$lib/stores/user.store";
    import Send from "svelte-material-icons/Send.svelte";
    import ImagePic from "svelte-material-icons/Image.svelte";
    import {ChatMessageModel} from "$lib/models/classes/chatMessage.model";
    import {CommentModel} from "$lib/models/classes/comment.model";
    import {chatMsgCat} from "$lib/models/enums/chatMsgCat";
    import {commentQueue} from "$lib/stores/commentQueue.store";
    import {replyComment} from "$lib/stores/replyComment.store";
    import ReplyComment from "$lib/components/chat/ReplyComment.svelte";
    import Spinner from "$lib/components/Spinner.svelte";
    import ImagePreviewGallery from "$lib/components/chat/ImagePreviewGallery.svelte";
    import {errorStore} from "$lib/stores/error.store";

    let textArea:HTMLTextAreaElement
    let message:string = ""
    let previewImage:string|ArrayBuffer
    let imageFiles:FileList
    let imageFile:File
    let fileName:string = ""
    let isDisabled = false
    let controler:AbortController

    function sendMsg() {
        message.trim()

        const comment = new ChatMessageModel<CommentModel>(
            chatMsgCat.COMMENT,
                new CommentModel(
                    message,
                    $currentUser.id,
                    $replyComment.createdAt != null ? $replyComment : null,
                    null,
                    true,
                    fileName
                )
        )
        console.log(comment)
        commentQueue.addComment(comment)
        resetTextArea()
    }

    function resetTextArea() {
        replyComment.close()
        message = ""
        textArea.style.height = "1.25rem"
        textArea.focus()
        fileName = ""
    }

    function sendOrResize(e:KeyboardEvent){
        textArea.style.height = "1px"
        textArea.style.height = `${textArea.scrollHeight}px`

        if(e.key == "Enter"){
            message = message.slice(0, -1)
            sendMsg()
        }
    }

    async function uploadImage(){
        let uploadableImage:Blob
        isDisabled = true
        imageFile = imageFiles[0]
        const reader = new FileReader();
        // Closure to capture the file information.
        reader.onload = (function(theFile) {
            return function(e) {
                previewImage = e.target.result
            };
        })(imageFile);
        // Read in the image file as a data URL.
        // if allowedExtensions = /(\.jpg|\.jpeg|\.png|\.gif)$/i

        reader.readAsDataURL(imageFile)

        fileName = Date.now() + "-" + imageFile.name

        const gifExtension = /(\.gif)$/i
        if (gifExtension.exec(fileName)){
            uploadableImage = imageFile
        } else {
            uploadableImage = await resizeImage(imageFile, 1500)
        }

        controler = new AbortController()
        const signal = controler.signal

        let fd = new FormData()
        fd.append('file', uploadableImage, fileName)
        const ok = await fetch("?/uploadChatImg", {method: "POST", body: fd, signal: signal}).then(() => isDisabled = false)
        if (!ok) {
            $errorStore = "Oops... something went wrong. Please try another file"
            cancelUpload()
        }
    }

    const resizeImage = (file:File, maxSize:number):Promise<Blob> => {
        const reader = new FileReader();
        const image = new Image();
        const canvas = document.createElement('canvas');
        const dataURItoBlob = (dataURI: string) => {
            const bytes = dataURI.split(',')[0].indexOf('base64') >= 0 ?
                atob(dataURI.split(',')[1]) :
                unescape(dataURI.split(',')[1]);
            const mime = dataURI.split(',')[0].split(':')[1].split(';')[0];
            const max = bytes.length;
            const ia = new Uint8Array(max);
            for (var i = 0; i < max; i++) ia[i] = bytes.charCodeAt(i);
            return new Blob([ia], {type:mime});
        };
        const resize = () => {
            let width = image.width;
            let height = image.height;

            if (width > height) {
                if (width > maxSize) {
                    height *= maxSize / width;
                    width = maxSize;
                }
            } else {
                if (height > maxSize) {
                    width *= maxSize / height;
                    height = maxSize;
                }
            }

            canvas.width = width;
            canvas.height = height;
            canvas.getContext('2d').drawImage(image, 0, 0, width, height);
            let dataUrl = canvas.toDataURL('image/jpeg');
            return dataURItoBlob(dataUrl);
        };

        return new Promise((ok, no) => {
            if (!file.type.match(/image.*/)) {
                no(new Error("Not an image"));
                return;
            }

            reader.onload = (readerEvent: any) => {
                image.onload = () => ok(resize());
                image.src = readerEvent.target.result;
            };
            reader.readAsDataURL(file);
        })
    };

    function cancelUpload() {
        imageFiles = null
        fileName = ''
        controler.abort()
    }

    $: if($replyComment) {
        if($replyComment.text !== undefined && textArea !== undefined) {
            textArea.focus()
        }
    }
</script>

<div>
    <div class="flex">
        <div class="flex flex-col-reverse">

            <label for="upload"
                   class="cursor-pointer text-typography-main"
                   on:change={uploadImage}>

                <div class="rounded-full bg-primary p-3">
                    <span class="flex">
                        <ImagePic size="1.4em"/>
                    </span>
                    <input id="upload" class="hidden" name="file" type="file" accept="image/*" bind:files={imageFiles}>
                </div>

            </label>

        </div>
        <div class="flex-1
                    mx-3
                    border-solid
                    bg-canvas-secondary
                    border-2
                    border-gray-400
                    p-2
                    rounded-lg
                    shadow-sm">

            <ReplyComment />
            <ImagePreviewGallery fileName={fileName} previewImage={previewImage} cancelUpload={cancelUpload}/>


            <textarea class="text-sm
                            h-5
                            p-0
                            overflow-hidden
                            border-0
                            focus:border-0
                            focus:outline-0"
                      name="msg"
                      bind:this={textArea}
                      bind:value={message}
                      on:keyup={e => sendOrResize(e)}></textarea>
        </div>

        <div class="flex flex-col-reverse">
            <button on:click={sendMsg} class="rounded-full py-3 {isDisabled}" disabled={isDisabled}>
                {#if isDisabled}
                    <Spinner size="sm" thickness="s" color="grey"/>
                {:else}
                    <Send size="1.4em"/>
                {/if}
            </button>
        </div>
    </div>
</div>

