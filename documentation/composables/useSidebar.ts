import { ref, provide, type InjectionKey, type Ref } from "vue";

export type UseSidebar = {
  isOpen: Ref<boolean>;
  toggleSidebar: () => void;
};

export const useSidebarKey = Symbol() as InjectionKey<UseSidebar>;

export default function useSidebar() {
  const isOpen = ref(true);

  const breakpoint = 1024;
  function toggleSidebar() {
    isOpen.value = !isOpen.value;
  }

  onMounted(() => {
    handleResize();
    window.addEventListener("resize", handleResize);
  });

  onUnmounted(() => {
    window.removeEventListener("resize", handleResize);
  });

  const handleResize = () => {
    if (window.innerWidth >= breakpoint) {
      // Desktop: always keep sidebar open
      isOpen.value = true;
    } else {
      // Mobile: close sidebar when switching to mobile view
      isOpen.value = false;
    }
  };

  provide(useSidebarKey, {
    isOpen,
    toggleSidebar,
  });
}
