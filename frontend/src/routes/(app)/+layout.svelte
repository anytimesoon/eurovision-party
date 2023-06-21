<script lang="ts">
  import {currentUser} from "$lib/stores/user.store";
  import {authLvl} from "$lib/models/enums/authLvl.enum";
  import {countryStore, participatingCountryStore} from "$lib/stores/country.store";
  import type {LayoutData} from "./$types";
  import {chatMsgCat} from "$lib/models/enums/chatMsgCat";
  import {CommentModel} from "$lib/models/classes/comment.model";
  import {commentStore} from "$lib/stores/comment.store";

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
    })


  };
</script>
  
<main>
  <nav>
    <a href="/">Chat</a>

    {#each $participatingCountryStore as country }
      <span><a href={"/country/" + country.slug}>{country.name}</a>{" "}</span>
    {/each}

    <a href="/results">Results</a>

    {#if $currentUser.authLvl === authLvl.ADMIN }
      <a href="/admin/countries">Countries</a>
      <a href="/admin/friends">Friends</a>
    {/if}

    <a href="/settings">Settings</a>
  </nav>
  
  <slot />
</main>