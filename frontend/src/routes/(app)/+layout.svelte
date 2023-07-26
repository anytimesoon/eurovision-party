<script lang="ts">
  import {currentUser} from "$lib/stores/user.store";
  import {authLvl} from "$lib/models/enums/authLvl.enum";
  import {countryStore, participatingCountryStore} from "$lib/stores/country.store";
  import type {LayoutData} from "./$types";
  import {chatMsgCat} from "$lib/models/enums/chatMsgCat";
  import {CommentModel} from "$lib/models/classes/comment.model";
  import {commentStore} from "$lib/stores/comment.store";
  import ChatIcon from "$lib/components/icons/ChatIcon.svelte";
  import VoteIcon from "$lib/components/icons/VoteIcon.svelte";
  import ResultIcon from "$lib/components/icons/ResultIcon.svelte";
  import SettingIcon from "$lib/components/icons/SettingIcon.svelte";
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

<!--  <nav class="flex flex-grow justify-center">-->
<!--    <a href="/">Chat</a>-->
<!--    <a href="#">Vote</a>-->
<!--    &lt;!&ndash;{#each $participatingCountryStore as country }&ndash;&gt;-->
<!--    &lt;!&ndash;  <span><a href={"/country/" + country.slug}>{country.name}</a>{" "}</span>&ndash;&gt;-->
<!--    &lt;!&ndash;{/each}&ndash;&gt;-->

<!--    <a href="/results">Results</a>-->

<!--    {#if $currentUser.authLvl === authLvl.ADMIN }-->
<!--      <a href="/admin/countries">Countries</a>-->
<!--      <a href="/admin/friends">Friends</a>-->
<!--    {/if}-->

<!--    <a href="/settings">Settings</a>-->
<!--  </nav>-->


</main>