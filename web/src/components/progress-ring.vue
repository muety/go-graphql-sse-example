<template>
  <div class="relative">
    <div class="absolute h-full w-full container items-center justify-center">
      <slot />
    </div>
    <svg
            class="progress-ring"
            :viewBox="viewBox"
    >
      <circle
              class="circle-foreground"
              :fill="fillColor"
              :stroke="color"
              :stroke-dasharray="circumference + ' ' + circumference"
              :style="{ strokeDashoffset }"
              :stroke-width="stroke"
              :r="normalizedRadius"
              :cx="radius"
              :cy="radius"
      />
      <circle
              class="circle-background"
              fill="transparent"
              stroke="#e7e5e4"
              :stroke-dasharray="circumference + ' ' + circumference"
              :style="{ strokeDashoffset: inverseStrokeDashoffset }"
              :stroke-width="stroke"
              :r="normalizedRadius"
              :cx="radius"
              :cy="radius"
      />
    </svg>
  </div>
</template>

<script>
export default {
  name: 'ProgressRing',
  props: {
    progress: {
      type: Number,
      required: true
    },
    stroke: {
      type: Number,
      required: false,
      default: 4
    },
    color: {
      type: String,
      required: false,
      default: '#42b983'
    },
    fillColor: {
      type: String,
      required: false,
      default: 'transparent'
    }
  },
  data () {
    return {
      radius: 100
    }
  },
  computed: {
    normalizedRadius () {
      return this.radius - this.stroke * 2
    },
    circumference () {
      return this.normalizedRadius * 2 * Math.PI
    },
    viewBox () {
      return `0 0 ${this.radius * 2} ${this.radius * 2}`
    },
    strokeDashoffset () {
      return this.circumference - this.progress / 100 * this.circumference
    },
    inverseStrokeDashoffset () {
      return this.circumference - (100 - this.progress) / 100 * this.circumference
    }
  }
}
</script>

<style scoped>
    .progress-ring {
        width: 100%;
        height: 100%;
    }

    .progress-ring .circle-foreground {
        transform: rotate(-90deg);
        transform-origin: 50% 50%;
    }

    .progress-ring .circle-background {
        transform-origin: 50% 50%;
        transform: scale(-1, 1) rotate(-90deg);
    }
</style>
