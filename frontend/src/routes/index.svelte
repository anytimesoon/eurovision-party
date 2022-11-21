<script lang="ts" context="module">
	import { UserModel } from "$lib/models/classes/user.model";
    import {CommentModel} from "src/lib/models/classes/comment.model"

    export let comments: CommentModel[] = [];
</script>

<script lang="ts">
	import { beforeUpdate, afterUpdate } from 'svelte';

	let div : any;
	let autoscroll : boolean;

	beforeUpdate(() => {
		// determine whether we should auto-scroll
		// once the DOM is updated...
		autoscroll = div && (div.offsetHeight + div.scrollTop) > (div.scrollHeight - 20);
	});

	afterUpdate(() => {
		// ...the DOM is now in sync with the data
		if (autoscroll) div.scrollTo(0, div.scrollHeight)
	});


	function handleKeydown(event : KeyboardEvent) {
		if (event.key === 'Enter') {
			const content : string = (event.target as HTMLInputElement).value;
			if (!content) return;

            let newComment : CommentModel = new CommentModel{
                userId = 'asdf',
                text = content
            }
			
            comments.push(newComment);

			(event.target as HTMLInputElement).value = '';

			setTimeout(() => {
                let incomming : CommentModel = new CommentModel {
					userId = 'alsdf',
					text = '...',
					placeholder = true
				}
				comments.push(incomming);

				setTimeout(() => {
                    let fullReply : CommentModel = new CommentModel() {
                        userId = 'asdfa',
                        text = 'reply'
                    }
					comments.filter(comment => !comment.placeholder).push(
						fullReply
					);
				}, 500 + Math.random() * 500);
			}, 200 + Math.random() * 200);
		}
	}
</script>

<style>
	.chat {
		display: flex;
		flex-direction: column;
		height: 100%;
		max-width: 320px;
	}

	.scrollable {
		flex: 1 1 auto;
		border-top: 1px solid #eee;
		margin: 0 0 0.5em 0;
		overflow-y: auto;
	}

	article {
		margin: 0.5em 0;
	}

	.user {
		text-align: right;
	}

	span {
		padding: 0.5em 1em;
		display: inline-block;
	}

	.eliza span {
		background-color: #eee;
		border-radius: 1em 1em 1em 0;
	}

	.user span {
		background-color: #0074D9;
		color: white;
		border-radius: 1em 1em 0 1em;
		word-break: break-all;
	}
</style>

<div class="chat">
	<h1>Eliza</h1>

	<div class="scrollable" bind:this={div}>
		{#each comments as comment}
			<article class={comment.userId}>
				<span>{comment.text}</span>
			</article>
		{/each}
	</div>

	<input on:keydown={handleKeydown}>
</div>
