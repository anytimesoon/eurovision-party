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
    import {staticEP} from "$lib/models/enums/endpoints.enum";
    import {sessionStore} from "$lib/stores/session.store";
    import {v4 as uuid} from 'uuid';

    let textArea:HTMLTextAreaElement = $state()
    let message:string = $state("")
    let previewImage:string|ArrayBuffer = $state()
    const authorizedExtensions = ['image/jpeg', 'image/jpg', 'image/png', 'image/gif', 'video/mp4', 'video/webm']
    let imageFiles:FileList = $state()
    let imageFile:File
    let fileName:string = $state("")
    let isDisabled = $state(false)
    let controller:AbortController

    function sendMsg() {
        const trimmedMessage = message.trim()
        if (trimmedMessage === "" && fileName === "") {
            $errorStore = "Say something! Messages can't be blank"
            textArea.focus()
            return
        }

        const comment = new ChatMessageModel<CommentModel>(
            chatMsgCat.COMMENT,
                new CommentModel(
                    trimmedMessage,
                    $currentUser.id,
                    uuid(),
                    $replyComment.createdAt != null ? $replyComment : null,
                    null,
                    true,
                    fileName
                )
        )

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
        isDisabled = true
        imageFile = imageFiles[0]
        controller = new AbortController()
        if (!authorizedExtensions.includes(imageFile.type)) {

            $errorStore = "Unsupported file type"
            cancelUpload()
            return
        }

        let uploadableImage:Blob
        const reader = new FileReader();
        // Closure to capture the file information.
        reader.onload = (function(theFile) {
            return function(e) {
                if(theFile.type.includes("image")) {
                    previewImage = e.target.result
                } else {
                    previewImage = staticEP.IMG + "video.png"
                }
            };
        })(imageFile);
        // Read in the image file as a data URL.
        reader.readAsDataURL(imageFile)

        fileName = Date.now() + "-" + $currentUser.id + imageFile.type.replace(/(image\/|video\/)/, ".")

        const gifExtension = /(\.gif|\.mp4|\.webm)$/i
        if (gifExtension.exec(fileName)){
            if (imageFile.size > 1024 * 1024 * 5) {
                $errorStore = "File is too large"
                cancelUpload()
                return
            }
            uploadableImage = imageFile
        } else {
            uploadableImage = await resizeImage(imageFile, 1500)
        }

        const signal = controller.signal

        let fd = new FormData()
        fd.append('file', uploadableImage, fileName)
        const resp = await fetch(staticEP.CREATE_CHAT_IMG, {
            method: "POST",
            body: fd,
            signal: signal,
            headers: {
                "Authorization": $sessionStore
            }
        })
        if (!resp.ok) {
            $errorStore = "Oops... something went wrong. Please try another file"
            cancelUpload()
        }
        isDisabled = false
    }

    const resizeImage = async (file:File, maxSize:number):Promise<Blob> => {
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
        controller.abort()
        isDisabled = false
    }

    $effect(() => {
        if($replyComment) {
            if($replyComment.text !== undefined && textArea !== undefined) {
                textArea.focus()
            }
        }
    });
</script>

<div>
    <div class="flex">
        <div class="flex flex-col-reverse">

            <label for="upload"
                   class="cursor-pointer text-typography-main"
                   onchange={uploadImage}>

                <span class="rounded-full bg-primary p-3 block">
                    <span class="flex">
                        <ImagePic size="1.4em"/>
                    </span>
                    <input id="upload" class="hidden" name="file" type="file" accept={authorizedExtensions.join(",")} bind:files={imageFiles}>
                </span>

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
                            border-none
                            bg-transparent
                            shadow-none
                            resize-none
                            outline-none
                            focus:border-none
                            focus:outline-none
                            focus:resize-none"
                      name="msg"
                      bind:this={textArea}
                      bind:value={message}
                      onkeyup={e => sendOrResize(e)}></textarea>
        </div>

        <div class="flex flex-col-reverse">
            <button onclick={sendMsg} class="rounded-full py-3 {isDisabled}" disabled={isDisabled}>
                {#if isDisabled}
                    <Spinner size="sm" thickness="s" color="grey"/>
                {:else}
                    <Send size="1.4em"/>
                {/if}
            </button>
        </div>
    </div>
</div>

