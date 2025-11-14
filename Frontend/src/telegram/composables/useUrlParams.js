import { ref } from 'vue'

/**
 * Композабл для извлечения параметров из URL
 * @returns {Object} Объект с реактивными переменными и методами
 */
export function useUrlParams() {
  const startParam = ref(null)

  /**
   * Извлечение параметров Telegram WebApp из URL
   */
  const extractTelegramParamsFromUrl = () => {
    try {
      // Получаем текущий URL
      const url = window.location.href;
      console.log('🔍 Текущий URL:', url);
      
      // Проверяем, содержит ли URL параметры Telegram WebApp после хэша
      const hashIndex = url.indexOf('#');
      if (hashIndex !== -1) {
        const hashPart = url.substring(hashIndex + 1);
        console.log('🔍 Часть URL после хэша:', hashPart);
        
        // Проверяем, содержит ли часть после хэша параметры Telegram WebApp
        if (hashPart.includes('tgWebAppData=') || hashPart.includes('tgWebAppStartParam=')) {
          console.log('🔍 Найдены параметры Telegram WebApp в URL');
          
          // Извлекаем параметры из части после хэша
          // Обрабатываем различные форматы URL
          let paramsString = '';
          if (hashPart.includes('?')) {
            // Формат: /#/path?param1=value1&param2=value2
            paramsString = hashPart.split('?')[1];
          } else if (hashPart.includes('&') || hashPart.includes('=')) {
            // Формат: /#/param1=value1&param2=value2
            paramsString = hashPart;
          } else {
            // Формат: /#/tgWebAppData=value1&tgWebAppStartParam=value2
            // Ищем начало параметров
            const tgWebAppDataIndex = hashPart.indexOf('tgWebAppData=');
            const tgWebAppStartParamIndex = hashPart.indexOf('tgWebAppStartParam=');
            
            if (tgWebAppDataIndex !== -1) {
              paramsString = hashPart.substring(tgWebAppDataIndex);
            } else if (tgWebAppStartParamIndex !== -1) {
              paramsString = hashPart.substring(tgWebAppStartParamIndex);
            }
          }
          
          if (paramsString) {
            try {
              const params = new URLSearchParams(paramsString);
              const tgWebAppData = params.get('tgWebAppData');
              const tgWebAppStartParam = params.get('tgWebAppStartParam');
              
              console.log('🔍 Извлеченные параметры:', { tgWebAppData, tgWebAppStartParam });
              
              // Если есть параметры, сохраняем их для дальнейшей обработки
              if (tgWebAppStartParam) {
                startParam.value = tgWebAppStartParam;
                console.log('✅ Сохранен параметр startParam из URL:', startParam.value);
              }
            } catch (parseError) {
              console.error('❌ Ошибка при парсинге параметров из URL:', parseError);
            }
          }
        }
      }
    } catch (error) {
      console.error('❌ Ошибка при извлечении параметров Telegram WebApp из URL:', error);
    }
  };

  return {
    startParam,
    extractTelegramParamsFromUrl
  }
}
