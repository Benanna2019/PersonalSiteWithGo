function PostPage({ html, state }) {
  const { store } = state;
  const { post } = store;
  const { frontmatter } = post;
  const { description = "", published = "", title = "" } = frontmatter;

  return html`
    <style>
      h1,
      .date {
        text-align: var(--align-heading);
      }
    </style>
    <site-layout>
      <article class="h-entry font-body leading4 mi-auto pb0 pb4-lg">
        <a href="/" class="text-link">← Back</a>

        <h1
          class="p-name font-heading font-bold mbe0 text4 tracking-1 leading1"
        >
          ${title}
        </h1>
        <p class="date dt-published mbe4">${published}</p>
        <section class="prose" slot="e-content doc">${post.html}</section>
        <section class="p-summary hidden">${description}</section>
      </article>
    </site-layout>
  `;
}

//   ${mentions?.length ? "<webmentions-list></webmentions-list>" : ""}
// <my-h-card class="hidden"></my-h-card>;
