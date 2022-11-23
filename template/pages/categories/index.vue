<template>
  <div>
    <TopBar />
    <NarBar :catelist="catelist"/>
    <Featured />
    <Category :catelist="catelist"/>
    <Footer />
  </div>
</template>

<script>
export default {
  async fetch() {
      this.catelist = await fetch(
        'http://localhost:8081/product'
      ).then(res => res.json())
    },
    scrollToTop: true
}
</script>

<style>

</style>
