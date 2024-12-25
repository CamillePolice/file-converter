<script setup>
import Header from './components/Header.vue'
import Navigation from './components/Navigation.vue'
import Hero from './components/sections/Hero.vue'
import About from './components/sections/About.vue'
import Experience from './components/sections/Experience.vue'
import Skills from './components/sections/Skills.vue'
import Projects from './components/sections/Projects.vue'
import Contact from './components/sections/Contact.vue'

const sections = ['about', 'experience', 'skills', 'projects', 'contact']
</script>

<template>
  <div class="relative">
    <Header />
    <Navigation :sections="sections" />
    <main class="scroll-smooth">
      <Hero />
      <About id="about" />
      <Experience id="experience" />
      <Skills id="skills" />
      <Projects id="projects" />
      <Contact id="contact" />
    </main>
  </div>
</template>

<style>
html {
  scroll-behavior: smooth;
}
</style>
