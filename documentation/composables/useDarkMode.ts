import type { InjectionKey, Ref } from 'vue';

export const useDarkModeKey = Symbol('darkMode') as InjectionKey<UseDarkMode>;

export interface UseDarkMode {
  isDark: Ref<boolean>;
  toggle: () => void;
}

export default function useDarkMode(): UseDarkMode {
  const isDark = useState('darkMode', () => false);

  onMounted(() => {
    const stored = localStorage.getItem('darkMode');
    if (stored !== null) {
      isDark.value = stored === 'true';
    } else {
      isDark.value = window.matchMedia('(prefers-color-scheme: dark)').matches;
    }

    updateDarkClass();

    window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', (e) => {
      if (localStorage.getItem('darkMode') === null) {
        isDark.value = e.matches;
      }
    });
  });

  watch(isDark, () => {
    updateDarkClass();
    if (typeof window !== 'undefined') {
      localStorage.setItem('darkMode', String(isDark.value));
    }
  });

  function updateDarkClass() {
    if (typeof document === 'undefined') {
        return;
    }
      if (isDark.value) {
        document.documentElement.classList.add('dark');
      } else {
        document.documentElement.classList.remove('dark');
      }
    
  }

  function toggle() {
    isDark.value = !isDark.value;
  }

  provide(useDarkModeKey, { isDark, toggle });

  return { isDark, toggle };
}
