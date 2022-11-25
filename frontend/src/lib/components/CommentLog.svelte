<script type="ts">
    import type { CommentModel } from "$lib/models/classes/comment.model";
    import { commentStore } from "$lib/stores/comment.store";
    import Comment from "$lib/components/Comment.svelte";
    let timeline : Array<CommentModel>;
    commentStore.subscribe(val => {
        timeline = val.comments
    })
</script>

<div>
    {#each timeline as comment}
        <Comment comment={comment} />
    {/each}
</div>
