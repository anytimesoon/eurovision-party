<script lang="ts">
  import {countryStore, participatingCountryStore} from "$lib/stores/country.store";
  import type {LayoutData} from "./$types";
  import {chatMsgCat} from "$lib/models/enums/chatMsgCat";
  import type {CommentModel} from "$lib/models/classes/comment.model";
  import {commentStore} from "$lib/stores/comment.store";
  import MenuButton from "$lib/components/buttons/MenuButton.svelte";
  import {onMount} from "svelte";


  export let data:LayoutData
  $countryStore = data.countries
  const socket = data.socket
  let menu:HTMLElement

  onMount(() => {
    menu = document.getElementById("menu")
  })

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
    menu.classList.remove('right-0')
    menu.classList.add('-right-56')
  }

  const handleWindowClick = (e:Event) => {
    if(menu.classList.contains("right-0") && !e.target.classList.contains("voteNav") && e.target !== menu){
      closeMenu()
    }
  };

</script>
<svelte:window on:click={handleWindowClick} />
<main class="h-screen max-w-screen-sm mx-auto p-3 relative">
  <div class="flex flex-col">
      <div id="content" class="flex h-[calc(100vh-5rem)] pb-3">
        <div class="flex-grow">
          <slot />
        </div>
      </div>
      <nav class="h-4 flex w-full flex-wrap items-center justify-between bg-[#FBFBFB] text-neutral-500 hover:text-neutral-700 focus:text-neutral-700 dark:bg-neutral-600">

            <MenuButton icon="chat" menu={menu}/>
            <MenuButton icon="votes" menu={menu}/>
            <MenuButton icon="results"  menu={menu}/>
            <MenuButton icon="settings"  menu={menu}/>

      </nav>
  </div>

  <aside id="menu" class="fixed top-0 -right-56 bg-white z-1 flex duration-500 h-screen overflow-auto">
    <div class="w-full flex flex-col p-5">
<!--      <a href="#" on:click={closeMenu} class="text-right text-4xl">&times;</a>-->
      <ul class="list-none">
        {#each $participatingCountryStore as country}
          <li class="py-2"><a href="/country/{country.slug}" on:click={closeMenu}>{country.flag} {country.name}</a></li>
        {/each}
      </ul>
    </div>
  </aside>

</main>











