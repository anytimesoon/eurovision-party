<script lang="ts">
  import {countryStore, participatingCountryStore} from "$lib/stores/country.store";
  import type {LayoutData} from "./$types";
  import {chatMsgCat} from "$lib/models/enums/chatMsgCat";
  import {CommentModel} from "$lib/models/classes/comment.model";
  import {commentStore} from "$lib/stores/comment.store";
  import MenuButton from "$lib/components/buttons/MenuButton.svelte";


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

  const closeMenu = () => {
    const menu = document.getElementById("menu")
    menu.classList.remove('w-56')
    menu.classList.add('w-0')
  }


</script>

<main class="flex">
  <div class="h-screen flex m-auto">
    <div class="flex-col flex">
      <div class="max-w-lg flex flex-grow">
        <div class="flex-grow grid grid-cols-1 overflow-auto">
          <slot />
        </div>
      </div>
      <nav class="flex w-full flex-wrap items-center justify-between bg-[#FBFBFB] p-2 text-neutral-500 hover:text-neutral-700 focus:text-neutral-700 dark:bg-neutral-600">
        <div class="flex w-full flex-wrap items-center justify-between px-3">
          <div class="flex-grow basis-auto items-center justify-center flex">

            <MenuButton icon="chat" />
            <MenuButton icon="votes"/>
            <MenuButton icon="results" />
            <MenuButton icon="settings" />

          </div>
        </div>
      </nav>
    </div>

  </div>


  <aside id="menu" class="w-0 z-1 overflow-auto flex duration-500">
    <div class="w-full 2flex flex-col p-5 space-y-4">
      <a href="#" on:click={closeMenu} class="text-right text-4xl">&times;</a>
      <ul class="list-none">
        {#each $participatingCountryStore as country}
          <li class="py-2"><a href="/country/{country.slug}">{country.flag} {country.name}</a></li>
        {/each}
      </ul>
    </div>
  </aside>
</main>











