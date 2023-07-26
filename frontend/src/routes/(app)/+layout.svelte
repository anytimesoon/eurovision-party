<script lang="ts">
  import {countryStore, participatingCountryStore} from "$lib/stores/country.store";
  import type {LayoutData} from "./$types";
  import {chatMsgCat} from "$lib/models/enums/chatMsgCat";
  import {CommentModel} from "$lib/models/classes/comment.model";
  import {commentStore} from "$lib/stores/comment.store";
  import MenuButton from "$lib/components/buttons/MenuButton.svelte";
  import { onMount } from "svelte";
  import {
    Offcanvas,
    Ripple,
    Dropdown,
    initTE,
  } from "tw-elements";

  onMount(() => {
    initTE({ Offcanvas, Ripple, Dropdown });
  });

  export let data:LayoutData
  $countryStore = data.countries
  const socket = data.socket

  socket.onmessage = function (event) {
    const split = event.data.split("\n")
    split.map((c:string)=>{
      const chatMessage = JSON.parse(c)
      switch (chatMessage.category) {
        case chatMsgCat.COMMENT:
          let comment:CommentModel = chatMessage.body
          comment.createdAt = new Date(chatMessage.body.createdAt)
                console.log(comment)
          commentStore.update(comments => {
            return [...comments, comment]
          });
          break
        case chatMsgCat.COMMENT_ARRAY:
          let commentModels:CommentModel[] = chatMessage.body
                console.log(commentModels)
          for (let i = 0; i < commentModels.length; i++) {
            commentModels[i].createdAt = new Date(commentModels[i].createdAt)
          }
          commentStore.update(comments => {
            return [...comments, ...commentModels]
          })
          break
        default:
          console.log("bad message: " + c)
      }

      let chatBox = document.getElementById("chat-box")
      console.log("scrolling")
      chatBox.scrollTop = chatBox.scrollHeight
    })


  };
</script>
  
<main class="min-h-screen">
  <div class="min-h-[91vh] max-h-[91vh] max-w-lg container mx-auto grid grid-cols-1">
    <slot />
  </div>

  <nav class="fixed bottom-0 flex w-full flex-wrap items-center justify-between bg-[#FBFBFB] p-2 text-neutral-500 shadow-lg hover:text-neutral-700 focus:text-neutral-700 dark:bg-neutral-600">
    <div class="flex w-full flex-wrap items-center justify-between px-3">
      <div class="flex-grow basis-auto items-center justify-center flex">

        <MenuButton icon="chat" />
        <MenuButton icon="votes" />
        <MenuButton icon="results" />
        <MenuButton icon="settings" />

      </div>
    </div>
  </nav>

  <div
          class="invisible fixed bottom-0 right-0 top-0 z-[1045] flex w-96 max-w-full translate-x-full flex-col border-none bg-white bg-clip-padding text-neutral-700 shadow-sm outline-none transition duration-300 ease-in-out dark:bg-neutral-800 dark:text-neutral-200 [&[data-te-offcanvas-show]]:transform-none"
          tabindex="-1"
          id="offcanvas"
          aria-labelledby="offcanvasRightLabel"
          data-te-offcanvas-init>
    <div class="flex items-center justify-between p-4">
      <h5
              class="mb-0 font-semibold leading-normal"
              id="offcanvasRightLabel">
        Offcanvas right
      </h5>
      <button
              type="button"
              class="box-content rounded-none border-none opacity-50 hover:no-underline hover:opacity-75 focus:opacity-100 focus:shadow-none focus:outline-none"
              data-te-offcanvas-dismiss>
      <span
              class="w-[1em] focus:opacity-100 disabled:pointer-events-none disabled:select-none disabled:opacity-25 [&.disabled]:pointer-events-none [&.disabled]:select-none [&.disabled]:opacity-25">
        <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke-width="1.5"
                stroke="currentColor"
                class="h-6 w-6">
          <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M6 18L18 6M6 6l12 12" />
        </svg>
      </span>
      </button>
    </div>
    <div class="offcanvas-body flex-grow overflow-y-auto p-4">...</div>
  </div>


</main>